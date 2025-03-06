// Loads pokemon data into SQLite from the pokemondb/database repo

package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/frasmataz/go-pokemon-db/controllers"
	"github.com/frasmataz/go-pokemon-db/models"

	"github.com/go-git/go-git/v5"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v2"
)

const checkoutDir = "data"

func main() {
	// Check out pokemon data if not checked out already
	_, err := os.Stat(checkoutDir)
	if os.IsNotExist(err) {
		log.Printf("Checking out pokemon data to ./%s", checkoutDir)
		_, err := git.PlainClone(checkoutDir, false, &git.CloneOptions{
			URL:      "https://github.com/pokemondb/database",
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatalf("Error checking out pokemon data: %v", err)
		}
	} else {
		log.Println("Pokemon data already checked out.")
	}

	// Load data from yaml
	_ = loadDataFile[map[string]models.Pokemon]("data/data/pokemon.yaml")
	_ = loadDataFile[map[string]models.PokemonForms]("data/data/pokemon-forms.yaml")

	// Open DB
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	controllers.CreateTables(db)
}

func loadDataFile[T any](path string) (out T) {
	pokemonFile, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error unmarshalling %s: %v", path, err)
	}

	err = yaml.Unmarshal(pokemonFile, &out)
	if err != nil {
		log.Fatalf("Error parsing %s: %v", path, err)
	}

	return out
}
