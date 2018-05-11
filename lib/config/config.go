// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017-2018 Vladimir "fat0troll" Hodakov

package config

import (
	"io/ioutil"
	"path/filepath"

	"bitbucket.org/pztrn/mogrus"
	"gopkg.in/yaml.v2"
)

// VERSION is the current bot's version
const VERSION = "0.7.4"

// DatabaseConnection handles database connection settings in config.yaml
type DatabaseConnection struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// TelegramConnection handles settings for Telegram connection in config.yaml
type TelegramConnection struct {
	APIToken      string `yaml:"api_token"`
	WebHookDomain string `yaml:"webhook_domain"`
	ListenAddress string `yaml:"listen_address"`
}

// ProxySettings handles settings for SOCKS5 proxy in config.yml
type ProxySettings struct {
	Enabled  bool   `yaml:"enabled"`
	Address  string `yaml:"address,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

// SpecialChats handles settings for special chats
type SpecialChats struct {
	AcademyID      string `yaml:"academy_id"`
	BastionID      string `yaml:"bastion_id"`
	DefaultID      string `yaml:"default_id"`
	HeadquartersID string `yaml:"headquarters_id"`
	GamesID        string `yaml:"games_id"`
}

// LoggingConfig handles log file configuration
type LoggingConfig struct {
	LogPath string `yaml:"log_path"`
}

// Config is a struct which represents config.yaml structure
type Config struct {
	Telegram     TelegramConnection `yaml:"telegram_connection"`
	Database     DatabaseConnection `yaml:"database_connection"`
	Proxy        ProxySettings      `yaml:"socks_proxy"`
	SpecialChats SpecialChats       `yaml:"special_chats"`
	Logs         LoggingConfig      `yaml:"logs"`
}

// Init is a configuration initializer
func (c *Config) Init(log *mogrus.LoggerHandler, configPath string) {
	fname, _ := filepath.Abs(configPath)
	yamlFile, yerr := ioutil.ReadFile(fname)
	if yerr != nil {
		log.Fatal("Can't read config file")
	} else {
		log.Info("Using " + configPath + " as config file.")
	}

	yperr := yaml.Unmarshal(yamlFile, c)
	if yperr != nil {
		log.Fatal("Can't parse config file")
	}
}

// New creates new empty Config object
func New() *Config {
	c := &Config{}
	return c
}
