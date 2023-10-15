package main

import (
	"log"

	"github.com/omegaatt36/gormigrate-demo/config"
	"github.com/omegaatt36/gormigrate-demo/database"
	"github.com/omegaatt36/gormigrate-demo/database/migration"
)

func rollback(mg *database.Migrator, rollbackTo string) error {
	if rollbackTo == "last" {
		return mg.Rollback()
	}

	return mg.RollbackTo(rollbackTo)
}

func main() {
	database.InitDb()

	rollbackTo := config.GetString("ROLLBACK_TO")

	mg := database.NewMigrator(database.Rdb,
		migration.InitModelList,
		migration.MigrationList)

	if rollbackTo != "" {
		if err := rollback(mg, rollbackTo); err != nil {
			log.Fatalf("Could not RollbackLast: %v", err)
		}
		return
	}

	err := mg.Upgrade()
	if err != nil {
		log.Fatalf("Could not Upgrade: %v", err)
	}

	log.Println("Migration did run successfully")
}
