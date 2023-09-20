package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

const schemaQuery = `
DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users 
(
    ID INTEGER PRIMARY KEY AUTOINCREMENT,
    firstName TEXT NULL,
	age INTEGER NULL,
	lastName TEXT NULL
);`

func New(db *sqlx.DB) (*Repository, error) {
	_, err := db.Exec(schemaQuery)
	if err != nil {
		return nil, err
	}
	return &Repository{db: db}, nil
}
