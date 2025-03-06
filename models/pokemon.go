package models

import (
	"database/sql"
	"log"
)

type Pokemon struct {
	NationalDexID int    `yaml:"national"`
	Name          string `yaml:"name"`
	Generation    int    `yaml:"gen"`
}

func (p Pokemon) CreateTable(db *sql.DB) {
	sqlStmt := `
	create table pokemon (
		id integer not null primary key,
		name text,
		generation integer
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error setting up pokemon table: %v", err)
	}
}
