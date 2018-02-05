// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package appcontext

import (
	"bitbucket.org/pztrn/flagger"
	"bitbucket.org/pztrn/mogrus"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/broadcaster/broadcasterinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/chatter/chatterinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/config"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/connections"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/datacache/datacacheinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/forwarder/forwarderinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/migrations/migrationsinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/orders/ordersinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/pinner/pinnerinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/pokedexer/pokedexerinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/reminder/reminderinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/router/routerinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/squader/squaderinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/statistics/statisticsinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/talkers/talkersinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/users/usersinterface"
	"git.wtfteam.pro/fat0troll/i2_bot/lib/welcomer/welcomerinterface"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
	"github.com/robfig/cron"
	"net/http"
	"os"
	"time"
)

// Context is an application context struct
type Context struct {
	StartupFlags *flagger.Flagger
	Cfg          *config.Config
	Cron         *cron.Cron
	Log          *mogrus.LoggerHandler
	Bot          *tgbotapi.BotAPI
	DataCache    datacacheinterface.DataCacheInterface
	Forwarder    forwarderinterface.ForwarderInterface
	Migrations   migrationsinterface.MigrationsInterface
	Router       routerinterface.RouterInterface
	Pokedexer    pokedexerinterface.PokedexerInterface
	Db           *sqlx.DB
	Talkers      talkersinterface.TalkersInterface
	Broadcaster  broadcasterinterface.BroadcasterInterface
	Welcomer     welcomerinterface.WelcomerInterface
	Pinner       pinnerinterface.PinnerInterface
	Reminder     reminderinterface.ReminderInterface
	Chatter      chatterinterface.ChatterInterface
	Squader      squaderinterface.SquaderInterface
	Users        usersinterface.UsersInterface
	Statistics   statisticsinterface.StatisticsInterface
	Orders       ordersinterface.OrdersInterface
}

// Init is a initialization function for context
func (c *Context) Init() {
	l := mogrus.New()
	l.Initialize()

	log := l.CreateLogger("i2_bot")
	log.CreateOutput("stdout", os.Stdout, true, "debug")
	c.Log = log

	c.StartupFlags = flagger.New(c.Log)
	c.StartupFlags.Initialize()

	// Adding available startup flags here
	configFlag := flagger.Flag{}
	configFlag.Name = "config"
	configFlag.Description = "Configuration file path"
	configFlag.Type = "string"
	configFlag.DefaultValue = "./config.yaml"
	err := c.StartupFlags.AddFlag(&configFlag)
	if err != nil {
		c.Log.Errorln(err)
	}
	c.StartupFlags.Parse()

	configPath, err := c.StartupFlags.GetStringValue("config")
	if err != nil {
		c.Log.Errorln(err)
		c.Log.Fatal("Can't get config file parameter from command line. Exiting.")
	}

	c.Cfg = config.New()
	c.Cfg.Init(c.Log, configPath)

	logFile, err := os.OpenFile(c.Cfg.Logs.LogPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	c.Log.CreateOutput("file="+c.Cfg.Logs.LogPath, logFile, true, "debug")

	c.Bot = connections.BotInit(c.Cfg, c.Log)
	c.Db = connections.DBInit(c.Cfg, c.Log)

	crontab := cron.New()
	c.Cron = crontab
}

// RegisterBroadcasterInterface registering broadcaster interface in application
func (c *Context) RegisterBroadcasterInterface(bi broadcasterinterface.BroadcasterInterface) {
	c.Broadcaster = bi
	c.Broadcaster.Init()
}

// RegisterChatterInterface registers chatter interface in application
func (c *Context) RegisterChatterInterface(ci chatterinterface.ChatterInterface) {
	c.Chatter = ci
	c.Chatter.Init()
}

// RegisterDataCacheInterface registers datacache interface in application
func (c *Context) RegisterDataCacheInterface(di datacacheinterface.DataCacheInterface) {
	c.DataCache = di
	c.DataCache.Init()
}

// RegisterForwarderInterface registers forwarder interface in application
func (c *Context) RegisterForwarderInterface(fi forwarderinterface.ForwarderInterface) {
	c.Forwarder = fi
	c.Forwarder.Init()
}

// RegisterMigrationsInterface registering migrations interface in application
func (c *Context) RegisterMigrationsInterface(mi migrationsinterface.MigrationsInterface) {
	c.Migrations = mi
	c.Migrations.Init()
}

// RegisterOrdersInterface registers orders interface in application
func (c *Context) RegisterOrdersInterface(oi ordersinterface.OrdersInterface) {
	c.Orders = oi
	c.Orders.Init()
}

// RegisterPinnerInterface registering pinner interface in application
func (c *Context) RegisterPinnerInterface(pi pinnerinterface.PinnerInterface) {
	c.Pinner = pi
	c.Pinner.Init()
}

// RegisterPokedexerInterface registering parsers interface in application
func (c *Context) RegisterPokedexerInterface(pi pokedexerinterface.PokedexerInterface) {
	c.Pokedexer = pi
}

// RegisterReminderInterface registering reminder interface in application
func (c *Context) RegisterReminderInterface(ri reminderinterface.ReminderInterface) {
	c.Reminder = ri
	c.Reminder.Init()
}

// RegisterRouterInterface registering router interface in application
func (c *Context) RegisterRouterInterface(ri routerinterface.RouterInterface) {
	c.Router = ri
	c.Router.Init()
}

// RegisterStatisticsInterface registers statistics interface in application
func (c *Context) RegisterStatisticsInterface(si statisticsinterface.StatisticsInterface) {
	c.Statistics = si
	c.Statistics.Init()
}

// RegisterSquaderInterface registers squader interface in application
func (c *Context) RegisterSquaderInterface(si squaderinterface.SquaderInterface) {
	c.Squader = si
	c.Squader.Init()
}

// RegisterTalkersInterface registering talkers interface in application
func (c *Context) RegisterTalkersInterface(ti talkersinterface.TalkersInterface) {
	c.Talkers = ti
	c.Talkers.Init()
}

// RegisterWelcomerInterface registering welcomer interface in application
func (c *Context) RegisterWelcomerInterface(wi welcomerinterface.WelcomerInterface) {
	c.Welcomer = wi
	c.Welcomer.Init()
}

// RegisterUsersInterface registers users interface in application
func (c *Context) RegisterUsersInterface(ui usersinterface.UsersInterface) {
	c.Users = ui
	c.Users.Init()
}

// RunDatabaseMigrations applies migrations on bot's startup
func (c *Context) RunDatabaseMigrations() {
	c.Migrations.SetDialect("mysql")
	c.Migrations.Migrate()
}

// StartBot starts listening for Telegram updates
func (c *Context) StartBot() {
	_, err := c.Bot.SetWebhook(tgbotapi.NewWebhook(c.Cfg.Telegram.WebHookDomain + c.Bot.Token))
	if err != nil {
		c.Log.Fatal(err.Error())
	}

	updates := c.Bot.ListenForWebhook("/" + c.Bot.Token)
	go http.ListenAndServe(c.Cfg.Telegram.ListenAddress, nil)

	c.Log.Info("Listening on " + c.Cfg.Telegram.ListenAddress)
	c.Log.Info("Webhook URL: " + c.Cfg.Telegram.WebHookDomain + c.Bot.Token)

	for update := range updates {
		if update.Message != nil {
			if update.Message.From != nil {
				if update.Message.Date > (int(time.Now().Unix()) - 5) {
					go c.Router.RouteRequest(&update)
				}
			}
		} else if update.InlineQuery != nil {
			c.Router.RouteInline(&update)
		} else if update.CallbackQuery != nil {
			c.Router.RouteCallback(&update)
		} else if update.ChosenInlineResult != nil {
			c.Log.Debug(update.ChosenInlineResult.ResultID)
		} else {
			continue
		}
	}
}
