package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // Injects ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
}
