package infra

import (
	"log"

	"github.com/DavidHuie/gomigrate"
)

type Migrator struct {
	*gomigrate.Migrator
}

func NewMigrator(db *RDB, migrationsPath string) (*Migrator, error) {
	migrator, err := gomigrate.NewMigrator(db.DB, gomigrate.Postgres{}, migrationsPath)

	if err != nil {
		return nil, err
	}

	return &Migrator{migrator}, nil
}

func (m Migrator) Run() error {
	err := m.Migrate()
	if err != nil {
		log.Printf("error during migration: %v\n", err)
		log.Printf("migration failed, trying to roll back...\n")

		err = m.Rollback()
		if err != nil {
			log.Fatalf("migration is seriously broken: %v\n", err)
		}
	}

	return err
}
