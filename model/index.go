package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var database *mongo.Database

func LoadModel(db *mongo.Database) {
	database = db
}
