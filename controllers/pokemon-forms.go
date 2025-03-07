package controllers

import (
	"database/sql"
	"log"
)

func CreatePokemonFormsTable(db *sql.DB) {
	sqlStmt := `
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
		log.Fatalf("Error setting up pokemon forms table: %v", err)
	}
}
