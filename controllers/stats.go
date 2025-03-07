package controllers

import (
	"database/sql"
	"log"
)

func CreateStatsTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS stats (
		id				INTEGER NOT NULL PRIMARY KEY,
		hp				INTEGER,
		attack			INTEGER,
		defence			INTEGER,
		special_attack	INTEGER,
		special_defence INTEGER
	);
	CREATE TABLE IF NOT EXISTS ev_yield (
		id				INTEGER NOT NULL PRIMARY KEY,
		hp				INTEGER,
		attack			INTEGER,
		defence			INTEGER,
		special_attack	INTEGER,
		special_defence INTEGER
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error setting up stats tables: %v", err)
	}
}
