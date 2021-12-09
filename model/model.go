package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Reddit struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" `
	Title  string             `json:"title" bson:"title,omitempty"`
	Board  string             `json:"board" bson:"board,omitempty"`
	Author *Author            `json:"author" bson:"author,omitempty"`
}

type Author struct {
	FirstName string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}
