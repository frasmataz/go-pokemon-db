package controllers

import (
	"database/sql"
	"log"

	"github.com/frasmataz/go-pokemon-db/models"
)

func CreatePokemonTable(db *sql.DB) {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS pokemon (
		id			INTEGER NOT NULL PRIMARY KEY,
		name		TEXT,
		generation	INTEGER
	);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("Error setting up pokemon table: %v", err)
	}
}

func SavePokemon(db *sql.DB, pokemon models.Pokemon) {
	sqlStmt := `
	INSERT INTO pokemon (
		id,
		name,
		generation
	)
	VALUES (
		?,
		?,
		?
	)
	`
	_, err := db.Exec(
		sqlStmt,
		pokemon.NationalDexID,
		pokemon.Name,
		pokemon.Generation,
	)
	if err != nil {
		log.Fatalf("Error saving pokemon: %v", err)
	}
}
