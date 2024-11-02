package main

import (
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func getAllFiles(db *gorm.DB, conf interface{}) []XmlFile {
	var XmlFiles []XmlFile

	db.Find(&XmlFiles, "active = ?", true)

	return XmlFiles
}
