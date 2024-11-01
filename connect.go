package main

import "database/sql"

func connect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db, err
}
