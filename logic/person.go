package logic

import (
	"time"

	"github.com/Luka-Spa/GoAPI/model"
	"github.com/Luka-Spa/GoAPI/repository"
)

type PersonLogic struct {
	repo repository.ICRUD[model.Person]
}

func NewPersonLogic(repository repository.ICRUD[model.Person]) *PersonLogic {
	return &PersonLogic{
		repo: repository,
	}
}

func (l *PersonLogic) All(qp repository.QueryParams) []model.Person {
	people, _ := l.repo.All(qp)
	return people
}

func (l *PersonLogic) ById(id string) (model.Person, error) {
	return l.repo.ById(id)
}

func (l *PersonLogic) Create(person model.Person) (model.Person, error) {
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()
	return l.repo.Create(person)
}

func (l *PersonLogic) Update(id string, person model.Person) (model.Person, error) {
	person.UpdatedAt = time.Now()
	return l.repo.Update(id, person)
}

func (l *PersonLogic) Delete(id string) error {
	return l.repo.Delete(id)
}
