package model

import (
	"fmt"
)

type Result struct {
	Id           string `bson:"_id"`
	Flag         string `bson:"flag"`
	Status       string `bson:"status"`
	TitleNameEng string `bson:"titleNameEng"`
	TitleNameTh  string `bson:"titleNameTh"`
}

type UserRepository interface {
	FindOne() (Result, error)
}

type MongoRepository struct {
	collection string
}

func (r *MongoRepository) FindOne() (Result, error) {
	fmt.Println("FindOne 55555555555555555")

	var result Result

	result.TitleNameTh = r.collection

	return result, nil
}
