// i2_bot – Instinct PokememBro Bot
// Copyright (c) 2017 Vladimir "fat0troll" Hodakov

package migrationsinterface

// MigrationsInterface implements Migrations for importing via appcontext.
type MigrationsInterface interface {
	Init()
	Migrate() error
	SetDialect(dialect string) error
}
