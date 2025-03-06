package controllers

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS pokemon (
		id			INTEGER NOT NULL PRIMARY KEY,
		name		TEXT,
		generation	INTEGER
	);
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
	CREATE TABLE IF NOT EXISTS pokemon_forms (
		name		TEXT NOT NULL PRIMARY KEY,
		release		TEXT,
		type1		TEXT,
		type2		TEXT,
		stats_id	INTEGER,
		species		TEXT,
		height		INTEGER,
		weight		INTEGER,
		gender		TEXT,
		catch_rate	INTEGER,
		base_exp	INTEGER,
		egg_cycles	INTEGER,
		friendship	TEXT,
		growth_rate TEXT,
		ev_yield_id	INTEGER,
		FOREIGN KEY (stats_id) REFERENCES stats (id),
		FOREIGN KEY (ev_yield_id) REFERENCES ev_yield (id)
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error setting up pokemon table: %v", err)
	}
}
