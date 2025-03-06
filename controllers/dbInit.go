package controllers

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE pokemon (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		generation INTEGER
	);
	CREATE TABLE stats (
		id INTEGER NOT NULL PRIMARY KEY,
		hp INTEGER,
		attack INTEGER,
		defence INTEGER,
		special_attack INTEGER,
		special_defence INTEGER
	);
	CREATE TABLE ev_yield (
		id INTEGER NOT NULL PRIMARY KEY,
		hp INTEGER,
		attack INTEGER,
		defence INTEGER,
		special_attack INTEGER,
		special_defence INTEGER
	);
	CREATE TABLE pokemon_forms (
		name		TEXT NOT NULL PRIMARY KEY,
		release		TEXT,
		type1		TEXT,
		type2		TEXT,
		FOREIGN KEY(stats) REFERENCES stats(id),
		species		TEXT,
		height		INTEGER,
		weight		INTEGER,
		gender		TEXT,
		catch_rate	INTEGER,
		base_exp	INTEGER,
		egg_cycles	INTEGER,
		friendship	TEXT,
		growth_rate TEXT,
		FOREIGN KEY(ev_yield) REFERENCES ev_yield(id)
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error setting up pokemon table: %v", err)
	}
}
