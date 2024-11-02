package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnect(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db, err
}

func DatabaseMigrate(db *gorm.DB) {
	fmt.Println("Migrating db...")

	err := db.AutoMigrate(&XmlFile{})
	if err != nil {
		panic(err)
	}
}
