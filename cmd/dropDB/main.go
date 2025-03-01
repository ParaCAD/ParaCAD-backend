package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Get the database connection details from the environment variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Connect to the database
	dbURL := fmt.Sprintf("postgresql://%s@%s/%s?connect_timeout=10&password=%s&sslmode=disable", dbUser, dbHost, dbName, dbPassword)
	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	// Drop the database
	fmt.Println("\033[196;435mDropping the database in 5 seconds...\033[0m")
	time.Sleep(5 * time.Second)

	_, err = db.Exec(`
	DROP TABLE IF EXISTS parameter_constraint_types CASCADE
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	DROP TABLE IF EXISTS template_parameters_constraints CASCADE
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	DROP TABLE IF EXISTS template_parameters CASCADE
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	DROP TABLE IF EXISTS templates CASCADE
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
	DROP TABLE IF EXISTS users CASCADE
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database dropped successfully!")
}
