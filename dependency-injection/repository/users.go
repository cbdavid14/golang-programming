package repository

import "context"

type User struct {
	ID       int64
	FistName string
	LastName string
	Age      int
}

const (
	qryInsertUser  = `INSERT INTO users (firstName, lastName) VALUES (?, ?)`
	qryGetUserById = `SELECT * FROM users WHERE id = ?`
	qryGetUsers    = `SELECT * FROM users`
)

func (r *Repository) GetUserByID(ctx context.Context, id int64) (*User, error) {
	u := &User{}
	err := r.db.GetContext(ctx, u, qryGetUserById, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *Repository) getUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	err := r.db.SelectContext(ctx, &users, qryGetUsers)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) SaveUser(ctx context.Context, FirstName string, LastName string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, FirstName, LastName)
	if err != nil {
		return err
	}
	return nil
}
