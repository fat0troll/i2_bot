// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacacheinterface

import (
	"github.com/fat0troll/i2_bot/lib/datamapping"
	"github.com/fat0troll/i2_bot/lib/dbmapping"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// DataCacheInterface implements DataCache for importing via appcontext.
type DataCacheInterface interface {
	Init()

	GetAllGroupChats() []dbmapping.Chat
	GetAllPrivateChats() []dbmapping.Chat
	GetChatByID(chatID int) (*dbmapping.Chat, error)
	GetOrCreateChat(update *tgbotapi.Update) (*dbmapping.Chat, error)
	GetGroupChatsByIDs(chatIDs []int) []dbmapping.Chat
	GetLeaguePrivateChats() []dbmapping.Chat
	UpdateChatTitle(chatID int, newTitle string) (*dbmapping.Chat, error)

	AddPlayerToSquad(relation *dbmapping.SquadPlayer) (int, error)
	GetAllSquadMembers(squadID int) []dbmapping.SquadPlayerFull
	GetAllSquadsChats() []dbmapping.Chat
	GetAllSquadsWithChats() []dbmapping.SquadChat
	GetAvailableSquadsChatsForUser(userID int) []dbmapping.Chat
	GetCommandersForSquad(squadID int) []dbmapping.Player
	GetSquadByID(squadID int) (*dbmapping.SquadChat, error)
	GetSquadByChatID(chatID int) (*dbmapping.Squad, error)
	GetSquadsChatsBySquadsIDs(squadsIDs []int) []dbmapping.Chat
	GetUserRoleInSquad(squadID int, playerID int) string
	GetUserRolesInSquads(userID int) []dbmapping.SquadPlayerFull

	AddPlayer(player *dbmapping.Player) (int, error)
	ChangePlayerKarma(addition int, playerID int) (*dbmapping.Player, error)
	GetOrCreatePlayerByTelegramID(telegramID int) (*dbmapping.Player, error)
	GetPlayerByID(playerID int) (*dbmapping.Player, error)
	GetPlayerByTelegramID(telegramID int) (*dbmapping.Player, error)
	UpdatePlayerFields(player *dbmapping.Player) (*dbmapping.Player, error)
	UpdatePlayerTimestamp(playerID int) error

	AddProfile(profile *dbmapping.Profile) (int, error)
	GetPlayersWithCurrentProfiles() map[int]*dbmapping.PlayerProfile
	GetProfileByID(profileID int) (*dbmapping.Profile, error)
	GetProfileByPlayerID(playerID int) (*dbmapping.Profile, error)

	AddTournamentReport(tournamentNumber int, playerID int, target string) (int, error)

	GetElementByID(elementID int) (*datamapping.Element, error)
	FindElementIDBySymbol(symbol string) (int, error)

	GetLeagueByID(leagueID int) (*datamapping.League, error)
	GetLeagueByEnglishName(name string) (*datamapping.League, error)
	GetLeagueByName(name string) (*datamapping.League, error)
	GetLeagueBySymbol(symbol string) (*datamapping.League, error)

	GetLevelByID(levelID int) (*datamapping.Level, error)

	FindLocationIDByName(name string) (int, error)

	GetAllPokememes() map[int]*datamapping.PokememeFull
	GetPokememeByID(pokememeID int) (*datamapping.PokememeFull, error)
	GetPokememeByName(name string) (*datamapping.PokememeFull, error)
	GetPokememesCountByGradeAndLocation(grade int, locationID int) int

	GetWeaponTypeByAttack(attack int) (*datamapping.Weapon, error)
	GetWeaponTypeByID(weaponID int) (*datamapping.Weapon, error)
	GetWeaponTypeByName(name string) (*datamapping.Weapon, error)
}
