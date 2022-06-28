package repository

import (
	"github.com/Luka-Spa/GoAPI/model"
)

type IPerson interface {
	All(qp QueryParams) ([]model.Person, error)
	Create(person model.Person) (model.Person, error)
	Update(id string, person model.Person) (model.Person, error)
	Delete(id string) error
}
