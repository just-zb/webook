package data

import (
	"github.com/just-zb/webook/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct {
	DB *gorm.DB
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
	}
	return &Data{DB: db}, cleanup, nil
}

func NewDB(debug bool, dataConf *Database) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataConf.Connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if debug {
		db.Logger.LogMode(1)
	}
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}
	return db, nil
}
