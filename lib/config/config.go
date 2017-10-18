// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package config

import (
	// stdlib
	"io/ioutil"
	"log"
	"path/filepath"
	// 3rd-party
	"gopkg.in/yaml.v2"
)

const VERSION = "0.29"

type DatabaseConnection struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type TelegramConnection struct {
	APIToken string `yaml:"api_token"`
}

type Config struct {
	Telegram TelegramConnection `yaml:"telegram_connection"`
	Database DatabaseConnection `yaml:"database_connection"`
}

func (c *Config) Init() {
	fname, _ := filepath.Abs("./config.yml")
	yamlFile, yerr := ioutil.ReadFile(fname)
	if yerr != nil {
		log.Fatal("Can't read config file")
	}

	yperr := yaml.Unmarshal(yamlFile, c)
	if yperr != nil {
		log.Fatal("Can't parse config file")
	}
}

func New() *Config {
	c := &Config{}
	return c
}
