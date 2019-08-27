package interfaces

import "github.com/jinzhu/gorm"

type IDatabaseConnection interface {
	GetDB() *gorm.DB
}

