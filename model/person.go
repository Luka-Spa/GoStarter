package model

import "time"

type Person struct {
	Id        string    `bson:"_id,omitempty" json:"id"`
	Firstname string    `bson:"first_name" json:"first_name" binding:"required"`
	Lastname  string    `bson:"last_name" json:"last_name" binding:"required"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
