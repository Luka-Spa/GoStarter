package repository

type IRepository interface {
	GetContext() interface{}
	Connect() error
	Disconnect() error
}
