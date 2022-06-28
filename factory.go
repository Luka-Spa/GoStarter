package main

import (
	"github.com/Luka-Spa/GoAPI/controller"
	"github.com/Luka-Spa/GoAPI/logic"
	"github.com/Luka-Spa/GoAPI/repository"
	"github.com/Luka-Spa/GoAPI/repository/mongo"
	mg "go.mongodb.org/mongo-driver/mongo"
)

var personRepository repository.IPerson
var personLogic *logic.PersonLogic
var rep repository.IRepository

func UseMongo() {
	rep = mongo.NewRepository()
	rep.Connect()
	personRepository = mongo.NewPersonRepository(rep.GetContext().(*mg.Database))
}

func InitLogic() {
	personLogic = logic.NewPersonLogic(personRepository)
}

func InitControllers() {
	controller := controller.NewHTTPRouter()
	controller.InitPerson(personLogic)
	controller.Run()
}
