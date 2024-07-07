package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	//Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}
}

func initializeDB() (*sql.DB, error) {
	//dbURL := os.Getenv("DATABASE_URL") // Make sure this environment variable is set in Render's settings
	//if dbURL == "" {
	//	log.Fatal("DATABASE_URL is not set")
	//}
	//
	//db, err := sql.Open("postgres", dbURL)
	//if err != nil {
	//	return nil, err
	//}
	//if err := db.Ping(); err != nil {
	//	return nil, err
	//}
	//
	//migrationUp(db)

	//return db, nil
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("user"), os.Getenv("password"), os.Getenv("host"),
		os.Getenv("port"), os.Getenv("dbname"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	migrationUp(db)

	return db, nil
}

func migrationUp(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///usr/src/app/migrations",
		"postgres", driver)
	//m, err := migrate.NewWithDatabaseInstance(
	//	"file://migrations",
	//	"postgres", driver)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// apply migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	fmt.Println("migrations up")
}
