package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DavidHuie/gomigrate"
)

// AddMigrations enables migrations
func (a *App) AddMigrations(path string) {
	migrator := a.buildMigrator()

	if a.missingMigrations(migrator) {
		a.Migrate(migrator)
		os.Exit(0)
	}
}

func (a *App) buildMigrator() gomigrate.Migrator {
	db := (*a.Database).Driver()

	m, err := gomigrate.NewMigrator(
		db.(*sql.DB),
		gomigrate.Postgres{},
		"./migrations",
	)

	if err != nil {
		log.Fatalln("migrations failed", err.Error())
	}

	return *m
}

func (a *App) missingMigrations(migrator gomigrate.Migrator) bool {
	migrations := migrator.Migrations(-1)
	missing := false

	for _, m := range migrations {
		if m.Status == gomigrate.Inactive {
			fmt.Println("Missing Migration: ", m.Name)
			missing = true
		}
	}

	return missing
}

// Migrate migrates database
func (a *App) Migrate(migrator gomigrate.Migrator) {
	if err := migrator.Migrate(); err != nil {
		log.Fatalln("migrations failed", err.Error())
	}
}

// Rollback rolsl
func (a *App) Rollback() {
	fmt.Println("rolling back database")

	migrator := a.buildMigrator()

	if err := migrator.Rollback(); err != nil {
		log.Fatalln("migrations failed", err.Error())
	}
}
