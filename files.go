package main

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func GetActiveFiles(db *gorm.DB) []XmlFile {
	var XmlFiles []XmlFile

	db.Find(&XmlFiles, "active = ?", true)

	return XmlFiles
}
