package model

import "github.com/google/uuid"

// User型を定義する
type User struct {
	ID       uuid.UUID
	name     string
	mail     string
	password string
	profile  string
	path     string
	point    int
}
