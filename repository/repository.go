package repository

type IRepository interface {
	GetContext() interface{}
	Connect() error
	Disconnect() error
}

type QueryParams struct {
	Limit int
	Page  int
}
