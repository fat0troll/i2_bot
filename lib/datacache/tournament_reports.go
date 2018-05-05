// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacache

import (
	"errors"
	"strconv"
	"time"

	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
)

func (dc *DataCache) initTournamentReports() {
	c.Log.Info("Initializing Tournament Reports storage...")
	dc.tournamentReports = make(map[int]*dbmapping.TournamentReport)
	dc.tournamentReportsByTournamentAndPlayer = make(map[int]map[int]*dbmapping.TournamentReport)
}

func (dc *DataCache) loadTournamentReports() {
	c.Log.Info("Load current TournamentReports data from database to DataCache...")
	tournamentReports := []dbmapping.TournamentReport{}
	err := c.Db.Select(&tournamentReports, "SELECT * FROM tournament_reports")
	if err != nil {
		// This is critical error and we need to stop immediately!
		c.Log.Fatal(err.Error())
	}

	dc.tournamentReportsMutex.Lock()
	for i := range tournamentReports {
		dc.tournamentReports[tournamentReports[i].ID] = &tournamentReports[i]
		if dc.tournamentReportsByTournamentAndPlayer[tournamentReports[i].TournamentNumber] == nil {
			dc.tournamentReportsByTournamentAndPlayer[tournamentReports[i].TournamentNumber] = make(map[int]*dbmapping.TournamentReport)
		}

		dc.tournamentReportsByTournamentAndPlayer[tournamentReports[i].TournamentNumber][tournamentReports[i].PlayerID] = &tournamentReports[i]
	}
	c.Log.Info("Loaded tournament reports in DataCache: " + strconv.Itoa(len(dc.tournamentReports)))
	dc.tournamentReportsMutex.Unlock()
}

// External functions

// AddTournamentReport creates new tournament report in database
func (dc *DataCache) AddTournamentReport(tournamentNumber int, playerID int, target string) (int, error) {
	if tournamentNumber < 3551 {
		return 0, errors.New("Can't save tournament from the very past")
	}
	if dc.players[playerID] == nil {
		return 0, errors.New("There is no player with ID = " + strconv.Itoa(playerID))
	}
	if target != "O" && target != "M" && target != "A" {
		return 0, errors.New("Malformed report")
	}

	dc.tournamentReportsMutex.Lock()
	if dc.tournamentReportsByTournamentAndPlayer[tournamentNumber] != nil {
		if dc.tournamentReportsByTournamentAndPlayer[tournamentNumber][playerID] != nil {
			dc.tournamentReportsMutex.Unlock()
			return dc.tournamentReportsByTournamentAndPlayer[tournamentNumber][playerID].ID, errors.New("There is already this report in database")
		}
	} else {
		dc.tournamentReportsByTournamentAndPlayer[tournamentNumber] = make(map[int]*dbmapping.TournamentReport)
	}
	dc.tournamentReportsMutex.Unlock()

	report := dbmapping.TournamentReport{}
	report.PlayerID = playerID
	report.TournamentNumber = tournamentNumber
	report.Target = target
	report.CreatedAt = time.Now().UTC()

	c.Log.Debug("Saving report...")
	_, err := c.Db.NamedExec("INSERT INTO tournament_reports VALUES(NULL, :player_id, :tournament_number, :target, :created_at)", &tournamentNumber)
	if err != nil {
		c.Log.Error(err.Error())
		return 0, err
	}

	newReport := dbmapping.TournamentReport{}
	err = c.Db.Get(&newReport, "SELECT * FROM tournament_reports WHERE player_id=? AND tournament_number=?", report.PlayerID, report.TournamentNumber)
	if err != nil {
		c.Log.Error(err.Error())
		return 0, err
	}

	dc.tournamentReportsMutex.Lock()
	dc.tournamentReports[newReport.ID] = &newReport
	dc.tournamentReportsByTournamentAndPlayer[tournamentNumber][playerID] = &newReport
	dc.tournamentReportsMutex.Unlock()

	dc.ChangePlayerKarma(3, newReport.PlayerID)

	return newReport.ID, nil
}
