package models

import (
	"time"

	"gorm.io/gorm"
)

// Model defines model. We need lower case ID.
type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"-"`
}
