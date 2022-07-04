package repository

type IRepository interface {
	GetContext() interface{}
	Connect() error
	Disconnect() error
}

type ICRUD[T interface{}] interface {
	All(qp QueryParams) ([]T, error)
	ById(id string) (T, error)
	Create(person T) (T, error)
	Update(id string, person T) (T, error)
	Delete(id string) error
}

type QueryParams struct {
	Limit int
	Page  int
}
