package mongo

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Luka-Spa/GoAPI/repository"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoRepository struct{}

var db *mongo.Client
var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func NewRepository() repository.IRepository {
	return &MongoRepository{}
}

func (m *MongoRepository) Connect() error {
	uri := os.Getenv("MONGO_URL")
	log.Info("Connecting to mongo..")
	db, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	defer m.Disconnect()
	var err error
	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		log.Errorln(err)
		return err
	}
	fmt.Println("Successfully connected to Mongo")
	return nil
}

func (*MongoRepository) GetContext() interface{} {
	return db
}

func (*MongoRepository) Disconnect() error {
	var err error
	if err = db.Disconnect(ctx); err != nil {
		log.Error(err)
	}
	return err
}
