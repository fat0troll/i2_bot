// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrations

import (
	"github.com/pressly/goose"
)

// Init adds all migrations to applications
func (m *Migrations) Init() {
	c.Log.Info("Initializing migrations...")
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
	goose.AddNamedMigration("15_create_chats.go", CreateChatsUp, CreateChatsDown)
	goose.AddNamedMigration("16_change_chat_type_column.go", ChangeChatTypeColumnUp, ChangeChatTypeColumnDown)
	goose.AddNamedMigration("17_change_profile_pokememes_columns.go", ChangeProfilePokememesColumnsUp, ChangeProfilePokememesColumnsDown)
	goose.AddNamedMigration("18_add_pokememes_wealth.go", AddPokememesWealthUp, AddPokememesWealthDown)
	goose.AddNamedMigration("19_create_broadcasts.go", CreateBroadcastsUp, CreateBroadcastsDown)
	goose.AddNamedMigration("20_create_squads.go", CreateSquadsUp, CreateSquadsDown)
	goose.AddNamedMigration("21_change_telegram_id_column.go", ChangeTelegramIDColumnUp, ChangeTelegramIDColumnDown)
	goose.AddNamedMigration("22_add_flood_chat_id.go", AddFloodChatIDUp, AddFloodChatIDDown)
	goose.AddNamedMigration("23_add_user_type.go", AddUserTypeUp, AddUserTypeDown)
	goose.AddNamedMigration("24_create_orders.go", CreateOrdersUp, CreateOrdersDown)
	goose.AddNamedMigration("25_remove_reusable.go", RemoveReusableUp, RemoveReusableDown)
	goose.AddNamedMigration("26_create_orders_completions.go", CreateOrdersCompletionsUp, CreateOrdersCompletionsDown)
	goose.AddNamedMigration("27_add_new_weapon.go", AddNewWeaponUp, AddNewWeaponDown)
	goose.AddNamedMigration("28_fix_locations.go", FixLocationsUp, FixLocationsDown)
	goose.AddNamedMigration("29_fix_leagues_names.go", FixLeaguesNamesUp, FixLeaguesNamesDown)
	goose.AddNamedMigration("30_create_alarms.go", CreateAlarmsUp, CreateAlarmsUp)
	goose.AddNamedMigration("31_change_squads_table.go", ChangeSquadsTableUp, ChangeSquadsTableDown)
	goose.AddNamedMigration("32_add_is_active_to_pokememes.go", AddIsActiveToPokememesUp, AddIsActiveToPokememesDown)
	goose.AddNamedMigration("33_delete_datamapped_tables.go", DeleteDataMappedTablesUp, DeleteDataMappedTablesDown)
	goose.AddNamedMigration("34_delete_pokememes_tables.go", DeletePokememesTablesUp, DeletePokememesTablesDown)
	goose.AddNamedMigration("35_add_karma_to_players.go", AddKarmaToPlayersUp, AddKarmaToPlayersDown)
	goose.AddNamedMigration("36_create_tournament_reports.go", CreateTournamentReportsUp, CreateTournamentReportsDown)
}

// Migrate migrates database to current version
func (m *Migrations) Migrate() error {
	c.Log.Printf("Starting database migrations...")
	err := goose.Up(c.Db.DB, ".")
	if err != nil {
		c.Log.Fatal(err)

		return err
	}

	return nil
}

// SetDialect sets dialect for migrations
func (m *Migrations) SetDialect(dialect string) error {
	return goose.SetDialect(dialect)
}
