// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2018 Vladimir "fat0troll" Hodakov

package datacacheinterface

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"source.wtfteam.pro/i2_bot/i2_bot/lib/dbmapping"
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
	GetOrCreatePlayerByTelegramID(telegramID int) (*dbmapping.Player, error)
	GetPlayerByID(playerID int) (*dbmapping.Player, error)
	GetPlayerByTelegramID(telegramID int) (*dbmapping.Player, error)
	UpdatePlayerFields(player *dbmapping.Player) (*dbmapping.Player, error)
	UpdatePlayerTimestamp(playerID int) error

	AddProfile(profile *dbmapping.Profile) (int, error)
	GetPlayersWithCurrentProfiles() map[int]*dbmapping.PlayerProfile
	GetProfileByID(profileID int) (*dbmapping.Profile, error)
	GetProfileByPlayerID(playerID int) (*dbmapping.Profile, error)

	AddPokememe(pokememeData map[string]string, pokememeLocations map[string]string, pokememeElements map[string]string) (int, error)
	GetAllPokememes() map[int]*dbmapping.PokememeFull
	GetPokememeByID(pokememeID int) (*dbmapping.PokememeFull, error)
	GetPokememeByName(name string) (*dbmapping.PokememeFull, error)
	DeletePokememeByID(pokememeID int) error
	UpdatePokememe(pokememeData map[string]string, pokememeLocations map[string]string, pokememeElements map[string]string) (int, error)

	GetLeagueBySymbol(symbol string) (*dbmapping.League, error)

	GetWeaponTypeByID(weaponID int) (*dbmapping.Weapon, error)
	GetWeaponTypeByName(name string) (*dbmapping.Weapon, error)
}
