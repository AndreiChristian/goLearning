package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DB *gorm.DB
}

func (c *Config) NewDB() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("test.dv"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
