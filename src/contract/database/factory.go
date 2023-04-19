package database

import (
	"gorm.io/gorm"
)

type Factory interface {
	Channel(channel string) (*gorm.DB, error)
	RegisterDb(channel string, dbResolver func() (*gorm.DB, error))
}
