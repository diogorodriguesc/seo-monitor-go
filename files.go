package main

import (
	_ "github.com/lib/pq"
)

func getAllFiles(conf interface{}) []XmlFile {
	db, err := connect(conf.(map[string]interface{})["parameters"].(map[interface{}]interface{})["postgres_connection_string"].(string))

	if err != nil {
		panic("failed to connect database")
	}

	var XmlFiles []XmlFile

	db.Find(&XmlFiles, "active = ?", true)

	return XmlFiles
}
