package database_test

import (
	"testing"

	"github.com/omegaatt36/gormigrate-demo/database"
	"github.com/omegaatt36/gormigrate-demo/database/migration"
	"github.com/stretchr/testify/assert"
)

func TestMigrate(t *testing.T) {
	s := assert.New(t)
	database.InitDb()
	mg := database.NewMigrator(
		database.Rdb,
		migration.InitModelList,
		migration.MigrationList)

	s.NoError(mg.Upgrade())
}
