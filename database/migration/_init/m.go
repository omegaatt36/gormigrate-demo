package _init

import "github.com/omegaatt36/gormigrate-demo/models"

// User defines user's db model.
type User struct {
	models.Model
	Email    string `gorm:"type:varchar(256);not null;unique"`
	Password string `gorm:"type:varchar(256);not null"`
	Name     string `gorm:"type:varchar(128);default:'';not null"`
	Disabled bool   `gorm:"default:false;not null"`
}

// ModelSchemaList v0 Model Structs
var ModelSchemaList = []any{
	&User{},
}
