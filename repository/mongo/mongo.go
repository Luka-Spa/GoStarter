package mongo

import (
	"context"
	"os"
	"time"

	"github.com/Luka-Spa/GoAPI/repository"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct{}

var client *mongo.Client
var db *mongo.Database

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

func NewRepository() repository.IRepository {
	return &mongoRepository{}
}

func (m *mongoRepository) Connect() error {
	uri := os.Getenv("MONGO_URL")
	log.Info("Connecting to mongo..")
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Errorln(err)
		return err
	}
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Errorln(err)
		return err
	}
	log.Infoln("Successfully connected to Mongo")
	db = client.Database(os.Getenv("MONGO_DATABASE"))
	return nil
}

func (*mongoRepository) GetContext() interface{} {
	return db
}

func (*mongoRepository) Disconnect() error {
	var err error
	log.Infoln("Closing mongo connection...")
	if err = client.Disconnect(ctx); err != nil {
		log.Infoln(err)
	}
	return err
}
