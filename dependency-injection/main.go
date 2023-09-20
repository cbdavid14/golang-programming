package main

import (
	"context"
	_ "github.com/mattn/go-sqlite3"
	"golang-programming/dependency-injection/configuration"
	"golang-programming/dependency-injection/database"
	"golang-programming/dependency-injection/repository"
)

func main() {
	var err error
	ctx := context.Background()
	// Load configuration
	config, err := configuration.Load("./dependency-injection/config.yaml")
	if err != nil {
		panic(err)
	}
	// Create database connection
	db, err := database.CreateSqliteConnection(ctx, config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo, err := repository.New(db)
	if err != nil {
		panic(err)
	}

	err = repo.SaveUser(ctx, "John", "Doe")
	if err != nil {
		panic(err)
	}
}
