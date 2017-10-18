// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	// stdlib
	"log"
	// 3rd-party
	"github.com/pressly/goose"
)

type Migrations struct{}

func (m *Migrations) Init() {
	log.Printf("Initializing migrations...")
	// All migrations are here
	goose.AddNamedMigration("1_hello.go", HelloUp, nil)
	goose.AddNamedMigration("2_create_players.go", CreatePlayersUp, CreatePlayersDown)
	goose.AddNamedMigration("3_create_profiles.go", CreateProfilesUp, CreateProfilesDown)
	goose.AddNamedMigration("4_create_pokememes.go", CreatePokememesUp, CreatePokememesDown)
	goose.AddNamedMigration("5_create_locations.go", CreateLocationsUp, CreateLocationsDown)
	goose.AddNamedMigration("6_create_elements.go", CreateElementsUp, CreateElementsDown)
	goose.AddNamedMigration("7_create_leagues.go", CreateLeaguesUp, CreateLeaguesDown)
	goose.AddNamedMigration("8_create_relations.go", CreateRelationsUp, CreateRelationsDown)
	goose.AddNamedMigration("9_update_locations.go", UpdateLocationsUp, UpdateLocationsDown)
	goose.AddNamedMigration("10_update_leagues.go", UpdateLeaguesUp, UpdateLeaguesDown)
	goose.AddNamedMigration("11_profile_data_additions.go", ProfileDataAdditionsUp, ProfileDataAdditionsDown)
	goose.AddNamedMigration("12_create_profile_relations.go", CreateProfileRelationsUp, CreateProfileRelationsDown)
	goose.AddNamedMigration("13_create_weapons_and_add_wealth.go", CreateWeaponsAndAddWealthUp, CreateWeaponsAndAddWealthDown)
	goose.AddNamedMigration("14_fix_time_element.go", FixTimeElementUp, FixTimeElementDown)
}

func (m *Migrations) Migrate() error {
	log.Printf("Starting database migrations...")
	err := goose.Up(c.Db.DB, ".")
	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

func (m *Migrations) SetDialect(dialect string) error {
	return goose.SetDialect(dialect)
}
