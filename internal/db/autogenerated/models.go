// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

type User struct {
	ID   int32  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
