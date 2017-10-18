// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrationsinterface

type MigrationsInterface interface {
	Init()
	Migrate() error
	SetDialect(dialect string) error
}
