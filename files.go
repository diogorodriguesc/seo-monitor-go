package main

import (
	_ "github.com/lib/pq"
)

type File struct {
	path string
}

func getAllFiles(conf interface{}) []File {
	db, err := connect(conf.(map[string]interface{})["parameters"].(map[interface{}]interface{})["postgres_connection_string"].(string))
	defer db.Close()

	rows, err := db.Query("SELECT * FROM files")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	files := []File{}

	for rows.Next() {
		var id int
		var path string
		if err := rows.Scan(&id, &path); err != nil {
			panic(err)
		}

		files = append(files, File{path: path})
	}

	return files
}
