package model

import "time"

type Person struct {
	Id        string
	Firstname string
	Lastname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
