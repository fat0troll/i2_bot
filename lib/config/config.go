// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"lab.pztrn.name/golibs/mogrus"
	"path/filepath"
)

// VERSION is the urrent bot's version
const VERSION = "0.6"

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
	APIToken string `yaml:"api_token"`
}

// SpecialChats handles settings for special chats
type SpecialChats struct {
	BastionID      string `yaml:"bastion_id"`
	DefaultID      string `yaml:"default_id"`
	HeadquartersID string `yaml:"headquarters_id"`
}

// LoggingConfig handles log file configuration
type LoggingConfig struct {
	LogPath string `yaml:"log_path"`
}

// Config is a struct which represents config.yaml structure
type Config struct {
	Telegram     TelegramConnection `yaml:"telegram_connection"`
	Database     DatabaseConnection `yaml:"database_connection"`
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
