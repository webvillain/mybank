package model

import (
	"context"
	"database/sql"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Database interface {
	ListUsers(ctx context.Context, db *sql.DB) ([]*User, error)
	SingleUser(ctx context.Context, db *sql.DB, id string) (*User, error)
	CreateNewUser(ctx context.Context, db *sql.DB, name string, email string) error
	UpdateUser(ctx context.Context, db *sql.DB, id string) error
	DeleteUser(ctx context.Context, db *sql.DB, id string) error
}
