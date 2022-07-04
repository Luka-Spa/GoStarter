package main

import (
	"github.com/Luka-Spa/GoAPI/controller"
	"github.com/Luka-Spa/GoAPI/logic"
	"github.com/Luka-Spa/GoAPI/model"
	"github.com/Luka-Spa/GoAPI/repository"
	"github.com/Luka-Spa/GoAPI/repository/mongo"
	mg "go.mongodb.org/mongo-driver/mongo"
)

var personRepository repository.ICRUD[model.Person]
var personLogic *logic.PersonLogic
var rep repository.IRepository

func UseMongo() {
	rep = mongo.NewRepository()
	rep.Connect()
	var db = rep.GetContext().(*mg.Database)
	personRepository = mongo.NewRepo[model.Person](db, "person")
}

func InitLogic() {
	personLogic = logic.NewPersonLogic(personRepository)
}

func InitControllers() {
	router := controller.NewHTTPRouter()
	router.InitPerson(personLogic)
	router.RunHTTPS()
}
