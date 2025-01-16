// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
)

type Product struct {
	ID              int32
	ProductUrl      string
	ProductCategory string
	Description     sql.NullString
	Price           string
	Status          int32
	Quantity        sql.NullInt32
	ProductName     string
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
}

type User struct {
	ID           int32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}
