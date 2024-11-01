package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&XmlFile{})

	return db, err
}
