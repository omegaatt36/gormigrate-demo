package migration

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/omegaatt36/gormigrate-demo/database/migration/_init"
	v1 "github.com/omegaatt36/gormigrate-demo/database/migration/v1"
)

// InitModelList is the model list to init db.
var InitModelList = _init.ModelSchemaList

// MigrationList is list of migrations.
var MigrationList = []*gormigrate.Migration{
	&v1.AddGender,
}
