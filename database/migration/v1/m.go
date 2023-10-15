package v1

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/omegaatt36/gormigrate-demo/models"
	"gorm.io/gorm"
)

// User defines user's db model.
type User struct {
	models.Model
	Gender string `gorm:"type:varchar(256);not null;unique"`
}

// AddGender adds column gender to table users.
var AddGender = gormigrate.Migration{
	ID: "2023-10-15-add-column-gender-to-table-users",
	Migrate: func(tx *gorm.DB) error {
		return tx.Migrator().AddColumn(&User{}, "gender")
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropColumn(&User{}, "gender")
	},
}
