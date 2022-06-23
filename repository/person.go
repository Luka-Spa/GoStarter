package repository

import (
	"github.com/Luka-Spa/GoAPI/model"
)

type IPerson interface {
	All() ([]model.Person, error)
	Create(person model.Person) (model.Person, error)
	Update(person model.Person) (model.Person, error)
	Delete(id string) error
}
