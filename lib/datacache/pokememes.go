// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"sort"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
	"strconv"
	"strings"
	"time"
)

func (dc *DataCache) initPokememes() {
	c.Log.Info("Initializing Pokememes storage...")
	dc.pokememes = make(map[int]*dbmapping.Pokememe)
	dc.fullPokememes = make(map[int]*dbmapping.PokememeFull)
}

func (dc *DataCache) loadPokememes() {
	c.Log.Info("Load current Pokememes data from database to DataCache...")
	pokememes := []dbmapping.Pokememe{}
	err := c.Db.Select(&pokememes, "SELECT * FROM pokememes WHERE is_active=1")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}
	pokememesElements := []dbmapping.PokememeElement{}
	err = c.Db.Select(&pokememesElements, "SELECT * FROM pokememes_elements")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err)
	}
	pokememesLocations := []dbmapping.PokememeLocation{}
	err = c.Db.Select(&pokememesLocations, "SELECT * FROM pokememes_locations")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err)
	}

	dc.pokememesMutex.Lock()
	dc.fullPokememesMutex.Lock()
	for i := range pokememes {
		dc.pokememes[pokememes[i].ID] = &pokememes[i]

		// Filling fullPokememes
		fullPokememe := dbmapping.PokememeFull{}
		elementsListed := []dbmapping.Element{}
		locationsListed := []dbmapping.Location{}

		for j := range pokememesLocations {
			if pokememesLocations[j].PokememeID == pokememes[i].ID {
				for l := range dc.locations {
					if pokememesLocations[j].LocationID == dc.locations[l].ID {
						locationsListed = append(locationsListed, *dc.locations[l])
					}
				}
			}
		}

		for k := range pokememesElements {
			if pokememesElements[k].PokememeID == pokememes[i].ID {
				for e := range dc.elements {
					if pokememesElements[k].ElementID == dc.elements[e].ID {
						elementsListed = append(elementsListed, *dc.elements[e])
					}
				}
			}
		}

		fullPokememe.Pokememe = pokememes[i]
		fullPokememe.Elements = elementsListed
		fullPokememe.Locations = locationsListed

		dc.fullPokememes[pokememes[i].ID] = &fullPokememe
	}
	c.Log.Info("Loaded pokememes with all additional information in DataCache: " + strconv.Itoa(len(dc.fullPokememes)))
	dc.pokememesMutex.Unlock()
	dc.fullPokememesMutex.Unlock()
}

// External functions

// AddPokememe adds pokememe from parser
func (dc *DataCache) AddPokememe(pokememeData map[string]string, pokememeLocations map[string]string, pokememeElements map[string]string) (int, error) {
	_, noerr := dc.GetPokememeByName(pokememeData["name"])
	if noerr == nil {
		return 0, errors.New("This pokememe already exists")
	}

	gradeInt := c.Statistics.GetPoints(pokememeData["grade"])
	attackInt := c.Statistics.GetPoints(pokememeData["attack"])
	hpInt := c.Statistics.GetPoints(pokememeData["hp"])
	mpInt := c.Statistics.GetPoints(pokememeData["mp"])
	defenceInt := attackInt
	if pokememeData["defence"] != "" {
		defenceInt = c.Statistics.GetPoints(pokememeData["defence"])
	}
	priceInt := c.Statistics.GetPoints(pokememeData["price"])
	creatorID := c.Statistics.GetPoints(pokememeData["creator_id"])

	if !(gradeInt != 0 && attackInt != 0 && hpInt != 0 && mpInt != 0 && defenceInt != 0 && priceInt != 0 && creatorID != 0) {
		return 0, errors.New("Some of the required numerical values are empty")
	}

	pokememe := dbmapping.Pokememe{}
	pokememe.Grade = gradeInt
	pokememe.Name = pokememeData["name"]
	pokememe.Description = pokememeData["description"]
	pokememe.Attack = attackInt
	pokememe.HP = hpInt
	pokememe.MP = mpInt
	pokememe.Defence = defenceInt
	pokememe.Price = priceInt
	if pokememeData["purchaseable"] == "true" {
		pokememe.Purchaseable = true
	} else {
		pokememe.Purchaseable = false
	}
	pokememe.ImageURL = pokememeData["image"]
	pokememe.PlayerID = creatorID
	pokememe.CreatedAt = time.Now().UTC()

	locations := []dbmapping.Location{}
	elements := []dbmapping.Element{}

	for i := range pokememeLocations {
		locationID, err := dc.findLocationIDByName(pokememeLocations[i])
		if err != nil {
			return 0, err
		}
		locations = append(locations, *dc.locations[locationID])
	}

	for i := range pokememeElements {
		elementID, err := dc.findElementIDBySymbol(pokememeElements[i])
		if err != nil {
			return 0, err
		}
		elements = append(elements, *dc.elements[elementID])
	}

	// All objects are prepared, let's fill database with it!
	c.Log.Debug("Filling pokememe...")
	_, err := c.Db.NamedExec("INSERT INTO pokememes VALUES(NULL, :grade, :name, :description, :attack, :hp, :mp, :defence, :price, :purchaseable, :image_url, :player_id, 1, :created_at)", &pokememe)
	if err != nil {
		return 0, err
	}

	c.Log.Debug("Finding newly added pokememe...")
	insertedPokememe := dbmapping.Pokememe{}
	err = c.Db.Get(&insertedPokememe, c.Db.Rebind("SELECT * FROM pokememes WHERE grade=? AND name=?"), pokememe.Grade, pokememe.Name)
	if err != nil {
		return 0, err
	}

	// Now we creating locations and elements links
	locationsAndElementsFilledSuccessfully := true
	c.Log.Debug("Filling locations...")
	for i := range locations {
		link := dbmapping.PokememeLocation{}
		link.PokememeID = insertedPokememe.ID
		link.LocationID = locations[i].ID
		link.CreatedAt = time.Now().UTC()

		_, err := c.Db.NamedExec("INSERT INTO pokememes_locations VALUES(NULL, :pokememe_id, :location_id, :created_at)", &link)
		if err != nil {
			c.Log.Error(err.Error())
			locationsAndElementsFilledSuccessfully = false
		}
	}

	c.Log.Debug("Filling elements...")
	for i := range elements {
		link := dbmapping.PokememeElement{}
		link.PokememeID = insertedPokememe.ID
		link.ElementID = elements[i].ID
		link.CreatedAt = time.Now().UTC()

		_, err := c.Db.NamedExec("INSERT INTO pokememes_elements VALUES(NULL, :pokememe_id, :element_id, :created_at)", &link)
		if err != nil {
			c.Log.Error(err.Error())
			locationsAndElementsFilledSuccessfully = false
		}
	}

	if !locationsAndElementsFilledSuccessfully {
		c.Log.Debug("All fucked up, removing what we have already added...")
		// There is something fucked up. In normal state we're should never reach this code
		_, err = c.Db.NamedExec("DELETE FROM pokememes_locations WHERE pokememe_id=:id", &insertedPokememe)
		if err != nil {
			c.Log.Error(err.Error())
		}
		_, err = c.Db.NamedExec("DELETE FROM pokememes_elements WHERE pokememe_id=:id", &insertedPokememe)
		if err != nil {
			c.Log.Error(err.Error())
		}
		_, err = c.Db.NamedExec("DELETE FROM pokememes where id=:id", &insertedPokememe)
		if err != nil {
			c.Log.Error(err.Error())
		}
		return 0, errors.New("Failed to add pokememe to database")
	}

	fullPokememe := dbmapping.PokememeFull{}
	fullPokememe.Pokememe = insertedPokememe
	fullPokememe.Locations = locations
	fullPokememe.Elements = elements

	// Filling data cache
	dc.pokememesMutex.Lock()
	dc.fullPokememesMutex.Lock()
	dc.pokememes[insertedPokememe.ID] = &insertedPokememe
	dc.fullPokememes[insertedPokememe.ID] = &fullPokememe
	dc.pokememesMutex.Unlock()
	dc.fullPokememesMutex.Unlock()

	return insertedPokememe.ID, nil
}

// GetAllPokememes returns all pokememes
// Index in resulted map counts all pokememes ordered by grade and alphabetically
func (dc *DataCache) GetAllPokememes() map[int]*dbmapping.PokememeFull {
	pokememes := make(map[int]*dbmapping.PokememeFull)
	dc.fullPokememesMutex.Lock()

	var keys []string
	keysToIDs := make(map[string]int)
	for i := range dc.fullPokememes {
		gradeKey := ""
		if dc.fullPokememes[i].Pokememe.Grade == 0 {
			gradeKey += "Z"
		} else {
			gradeKey += string(rune('A' - 1 + dc.fullPokememes[i].Pokememe.Grade))
		}
		keys = append(keys, gradeKey+"_"+strconv.Itoa(dc.fullPokememes[i].Pokememe.Attack+100000000000000)+"_"+dc.fullPokememes[i].Pokememe.Name)
		keysToIDs[gradeKey+"_"+strconv.Itoa(dc.fullPokememes[i].Pokememe.Attack+100000000000000)+"_"+dc.fullPokememes[i].Pokememe.Name] = i
	}
	sort.Strings(keys)

	idx := 0
	for _, k := range keys {
		pokememes[idx] = dc.fullPokememes[keysToIDs[k]]
		idx++
	}

	dc.fullPokememesMutex.Unlock()
	return pokememes
}

// GetPokememeByID returns pokememe with additional information by ID
func (dc *DataCache) GetPokememeByID(pokememeID int) (*dbmapping.PokememeFull, error) {
	dc.fullPokememesMutex.Lock()
	if dc.fullPokememes[pokememeID] != nil {
		dc.fullPokememesMutex.Unlock()
		return dc.fullPokememes[pokememeID], nil
	}
	dc.fullPokememesMutex.Unlock()
	return nil, errors.New("There is no pokememe with ID = " + strconv.Itoa(pokememeID))
}

// GetPokememeByName returns pokememe from datacache by name
func (dc *DataCache) GetPokememeByName(name string) (*dbmapping.PokememeFull, error) {
	dc.fullPokememesMutex.Lock()
	for i := range dc.fullPokememes {
		if strings.HasPrefix(dc.fullPokememes[i].Pokememe.Name, name) {
			dc.fullPokememesMutex.Unlock()
			return dc.fullPokememes[i], nil
		}
	}
	dc.fullPokememesMutex.Unlock()
	return nil, errors.New("There is no pokememe with name = " + name)
}

// DeletePokememeByID removes pokememe from database
func (dc *DataCache) DeletePokememeByID(pokememeID int) error {
	pokememe, err := dc.GetPokememeByID(pokememeID)
	if err != nil {
		return err
	}

	_, err = c.Db.NamedExec("DELETE FROM pokememes_locations WHERE pokememe_id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Error(err.Error())
	}
	_, err = c.Db.NamedExec("DELETE FROM pokememes_elements WHERE pokememe_id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Error(err.Error())
	}
	_, err = c.Db.NamedExec("DELETE FROM pokememes where id=:id", &pokememe.Pokememe)
	if err != nil {
		c.Log.Error(err.Error())
	}

	dc.pokememesMutex.Lock()
	dc.fullPokememesMutex.Lock()
	delete(dc.pokememes, pokememe.Pokememe.ID)
	delete(dc.fullPokememes, pokememe.Pokememe.ID)
	dc.pokememesMutex.Unlock()
	dc.fullPokememesMutex.Unlock()
	return nil
}

// UpdatePokememe updates existing pokememes in database and datacache
func (dc *DataCache) UpdatePokememe(pokememeData map[string]string, pokememeLocations map[string]string, pokememeElements map[string]string) (int, error) {
	knownPokememe, err := dc.GetPokememeByName(pokememeData["name"])
	if err != nil {
		// This should never happen, but who knows?
		return 0, errors.New("This pokememe doesn't exist. We should add it instead")
	}

	gradeInt := c.Statistics.GetPoints(pokememeData["grade"])
	attackInt := c.Statistics.GetPoints(pokememeData["attack"])
	hpInt := c.Statistics.GetPoints(pokememeData["hp"])
	mpInt := c.Statistics.GetPoints(pokememeData["mp"])
	defenceInt := attackInt
	if pokememeData["defence"] != "" {
		defenceInt = c.Statistics.GetPoints(pokememeData["defence"])
	}
	priceInt := c.Statistics.GetPoints(pokememeData["price"])
	creatorID := c.Statistics.GetPoints(pokememeData["creator_id"])

	if !(gradeInt != 0 && attackInt != 0 && hpInt != 0 && mpInt != 0 && defenceInt != 0 && priceInt != 0 && creatorID != 0) {
		return 0, errors.New("Some of the required numerical values are empty")
	}

	pokememe := knownPokememe.Pokememe
	pokememe.Grade = gradeInt
	pokememe.Name = pokememeData["name"]
	pokememe.Description = pokememeData["description"]
	pokememe.Attack = attackInt
	pokememe.HP = hpInt
	pokememe.MP = mpInt
	pokememe.Defence = defenceInt
	pokememe.Price = priceInt
	if pokememeData["purchaseable"] == "true" {
		pokememe.Purchaseable = true
	} else {
		pokememe.Purchaseable = false
	}
	pokememe.ImageURL = pokememeData["image"]
	pokememe.PlayerID = creatorID
	pokememe.CreatedAt = time.Now().UTC()

	locations := []dbmapping.Location{}
	elements := []dbmapping.Element{}

	for i := range pokememeLocations {
		locationID, err := dc.findLocationIDByName(pokememeLocations[i])
		if err != nil {
			return 0, err
		}
		locations = append(locations, *dc.locations[locationID])
	}

	for i := range pokememeElements {
		elementID, err := dc.findElementIDBySymbol(pokememeElements[i])
		if err != nil {
			return 0, err
		}
		elements = append(elements, *dc.elements[elementID])
	}

	// All objects are prepared, let's fill database with it!
	c.Log.Debug("Updating existing pokememe...")
	_, err = c.Db.NamedExec("UPDATE pokememes SET grade=:grade, name=:name, description=:description, attack=:attack, hp=:hp, mp=:mp, defence=:defence, price=:price, purchaseable=:purchaseable, image_url=:image_url, player_id=:player_id, created_at=:created_at WHERE id=:id", &pokememe)
	if err != nil {
		return 0, err
	}

	// Now we creating locations and elements links
	locationsAndElementsFilledSuccessfully := true
	c.Log.Debug("Destroying old relations...")
	_, err = c.Db.NamedExec("DELETE FROM pokememes_locations WHERE pokememe_id=:id", &pokememe)
	if err != nil {
		c.Log.Error(err.Error())
		locationsAndElementsFilledSuccessfully = false
	}
	_, err = c.Db.NamedExec("DELETE FROM pokememes_elements WHERE pokememe_id=:id", &pokememe)
	if err != nil {
		c.Log.Error(err.Error())
		locationsAndElementsFilledSuccessfully = false
	}
	c.Log.Debug("Filling locations...")
	for i := range locations {
		link := dbmapping.PokememeLocation{}
		link.PokememeID = pokememe.ID
		link.LocationID = locations[i].ID
		link.CreatedAt = time.Now().UTC()

		_, err := c.Db.NamedExec("INSERT INTO pokememes_locations VALUES(NULL, :pokememe_id, :location_id, :created_at)", &link)
		if err != nil {
			c.Log.Error(err.Error())
			locationsAndElementsFilledSuccessfully = false
		}
	}

	c.Log.Debug("Filling elements...")
	for i := range elements {
		link := dbmapping.PokememeElement{}
		link.PokememeID = pokememe.ID
		link.ElementID = elements[i].ID
		link.CreatedAt = time.Now().UTC()

		_, err := c.Db.NamedExec("INSERT INTO pokememes_elements VALUES(NULL, :pokememe_id, :element_id, :created_at)", &link)
		if err != nil {
			c.Log.Error(err.Error())
			locationsAndElementsFilledSuccessfully = false
		}
	}

	if !locationsAndElementsFilledSuccessfully {
		c.Log.Debug("All fucked up, removing what we have already added...")
		// There is something fucked up. In normal state we're should never reach this code
		_, err = c.Db.NamedExec("DELETE FROM pokememes_locations WHERE pokememe_id=:id", &pokememe)
		if err != nil {
			c.Log.Error(err.Error())
		}
		_, err = c.Db.NamedExec("DELETE FROM pokememes_elements WHERE pokememe_id=:id", &pokememe)
		if err != nil {
			c.Log.Error(err.Error())
		}
		_, err = c.Db.NamedExec("DELETE FROM pokememes where id=:id", &pokememe)
		if err != nil {
			c.Log.Error(err.Error())
		}
		return 0, errors.New("Failed to add pokememe to database")
	}

	fullPokememe := dbmapping.PokememeFull{}
	fullPokememe.Pokememe = pokememe
	fullPokememe.Locations = locations
	fullPokememe.Elements = elements

	// Filling data cache
	dc.pokememesMutex.Lock()
	dc.fullPokememesMutex.Lock()
	dc.pokememes[pokememe.ID] = &pokememe
	dc.fullPokememes[pokememe.ID] = &fullPokememe
	dc.pokememesMutex.Unlock()
	dc.fullPokememesMutex.Unlock()

	return pokememe.ID, nil
}
