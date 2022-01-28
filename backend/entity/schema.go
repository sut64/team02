package entity

import (
	"gorm.io/gorm"

	"time"

)

type BookType struct {
	gorm.Model

	Type string `gorm:"uniqueIndex"`

	BookInformation []BookInformation `gorm:"foreignKey:BookTypeID"`
}

