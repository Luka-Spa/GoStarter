package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Luka-Spa/GoAPI/model"
	"github.com/Luka-Spa/GoAPI/repository"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrInvalidPersonID = errors.New("personID should be 6 length string")
	ErrPersonNotFound  = errors.New("person not found")
)

type personRepository struct {
	person *mongo.Collection
}

func NewPersonRepository(db *mongo.Database) repository.IPerson {
	return &personRepository{person: db.Collection("person")}
}

func (r *personRepository) All() ([]model.Person, error) {
	var people []model.Person
	result, err := r.person.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Errorln(err)
	}
	if err = result.All(context.TODO(), &people); err != nil {
		log.Errorln(err)
	}
	return people, err
}

func (r *personRepository) Create(person model.Person) (model.Person, error) {
	person.CreatedAt = time.Now()
	person.UpdatedAt = time.Now()
	result, err := r.person.InsertOne(context.TODO(), person)
	if err != nil {
		log.Errorln(result, err)
	}
	person.Id = result.InsertedID.(primitive.ObjectID).Hex()
	return person, err
}

func (r *personRepository) Update(id string, person model.Person) (model.Person, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorln(err)
		return model.Person{}, err
	}
	var update = bson.M{"$set": bson.M{"first_name": person.Firstname, "last_name": person.Lastname, "updated_at": time.Now()}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	res := r.person.FindOneAndUpdate(context.TODO(), bson.M{"_id": _id}, update, &opt)
	if res.Err() != nil {
		log.Errorln(res.Err())
	}
	var updated model.Person
	res.Decode(&updated)
	return updated, err
}
func (r *personRepository) Delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorln(err)
		return err
	}
	result, err := r.person.DeleteOne(context.TODO(), bson.M{"_id": _id})
	if err != nil {
		log.Errorln(err)
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("not found")
	}
	return nil
}
