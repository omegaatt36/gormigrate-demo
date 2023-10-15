package database

import (
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var migrationOptions = gormigrate.Options{
	UseTransaction: true,
}

// Migrator runs migration.
type Migrator struct {
	mg *gormigrate.Gormigrate
	// db         *gorm.DB
	models     []any
	migrations []*gormigrate.Migration
}

// NewMigrator creates migrator.
func NewMigrator(db *gorm.DB, initModels []any, migrations []*gormigrate.Migration) *Migrator {
	return &Migrator{
		mg:         gormigrate.New(db, &migrationOptions, migrations),
		models:     initModels,
		migrations: migrations,
	}
}

func (m *Migrator) initMigrate() error {
	m.mg.InitSchema(func(tx *gorm.DB) error {
		if err := tx.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
			return errors.Wrapf(err, "install UUID")
		}

		err := tx.AutoMigrate(m.models...)
		if err != nil {
			return err
		}
		return nil
	})

	return nil
}

// Upgrade upgrades db schema version.
func (m *Migrator) Upgrade() error {
	if err := m.initMigrate(); err != nil {
		return errors.Wrap(err, "initMigrate")
	}

	if len(m.migrations) == 0 {
		log.Printf("no migration found")
		return nil
	}

	if err := m.mg.Migrate(); err != nil {
		return errors.Wrap(err, "upgradeLatestMigrate")
	}

	log.Printf("upgraded to version \"%s\"", m.migrations[len(m.migrations)-1].ID)
	return nil
}

// Rollback rollbacks the last migration.
func (m *Migrator) Rollback() error {
	if err := m.mg.RollbackLast(); err != nil {
		return errors.Wrap(err, "RollbackLast")
	}

	log.Print("rollback to last")
	return nil
}

// Rollback rollbacks the last migration.
func (m *Migrator) RollbackTo(versionID string) error {
	if err := m.mg.RollbackTo(versionID); err != nil {
		return errors.Wrapf(err, "RollbackTo %s", versionID)
	}

	log.Printf("rollback to %s", versionID)
	return nil
}
