package repository

import "context"

type User struct {
	ID       int64  `db:"ID"`
	FistName string `db:"firstName"`
	LastName string `db:"lastName"`
	Age      int    `db:"age"`
}

const (
	qryInsertUser  = `INSERT INTO users (firstName, lastName, age) VALUES (?, ?, ?)`
	qryGetUserById = `SELECT * FROM users WHERE ID = ?`
	qryGetUsers    = `SELECT * FROM users`
)

func (r *Repository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	var u User
	err := r.db.GetContext(ctx, &u, qryGetUserById, id)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	err := r.db.SelectContext(ctx, &users, qryGetUsers)
	if err != nil {
		return users, err
	}
	return users, nil
}
func (r *Repository) SaveUser(ctx context.Context, firstName string, lastName string, age int64) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, firstName, lastName, age)
	if err != nil {
		return err
	}
	return nil
}
