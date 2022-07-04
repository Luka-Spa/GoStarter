package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/Luka-Spa/GoAPI/repository"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type repo[T interface{}] struct {
	c *mongo.Collection
}

func NewRepo[T interface{}](db *mongo.Database, collection string) repository.ICRUD[T] {
	return &repo[T]{c: db.Collection(collection)}
}
func (r *repo[T]) All(qp repository.QueryParams) ([]T, error) {
	var people []T
	var ops = options.Find()
	skip := qp.Page * qp.Limit
	if skip > -1 {
		ops.SetSkip(int64(skip))
	}
	if qp.Limit > -1 {
		ops.SetLimit(int64(qp.Limit))
	}
	result, err := r.c.Find(context.TODO(), bson.D{}, ops)
	if err != nil {
		log.Errorln(err)
	}
	if err = result.All(context.TODO(), &people); err != nil {
		log.Errorln(err)
	}
	return people, err
}

func (r *repo[T]) ById(id string) (T, error) {
	var obj T
	var err error
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorln(err)
		var errObj T
		return errObj, err
	}
	result := r.c.FindOne(context.TODO(), bson.M{"_id": _id})
	if err = result.Decode(&obj); err != nil {
		log.Errorln(err)
	}
	return obj, err
}

func (r *repo[T]) Create(obj T) (T, error) {
	var upsert = true
	var after = options.After
	var ops = &options.FindOneAndUpdateOptions{Upsert: &upsert, ReturnDocument: &after}
	var inserted T
	var err error
	result := r.c.FindOneAndUpdate(context.TODO(), bson.M{"_id": primitive.NewObjectIDFromTimestamp(time.Now())}, bson.M{"$set": obj}, ops)
	if err = result.Err(); err != nil {
		log.Errorln(result, result.Err())
	}
	result.Decode(&inserted)
	return inserted, err
}

func (r *repo[T]) Update(id string, obj T) (T, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorln(err)
		var errObj T
		return errObj, err
	}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	res := r.c.FindOneAndUpdate(context.TODO(), bson.M{"_id": _id}, bson.M{"$set": obj}, &opt)
	if res.Err() != nil {
		log.Errorln(res.Err())
	}
	var updated T
	res.Decode(&updated)
	return updated, err
}
func (r *repo[T]) Delete(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorln(err)
		return err
	}
	result, err := r.c.DeleteOne(context.TODO(), bson.M{"_id": _id})
	if err != nil {
		log.Errorln(err)
		return err
	}
	if result.DeletedCount < 1 {
		return errors.New("not found")
	}
	return nil
}
