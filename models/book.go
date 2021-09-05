package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	Name   string             `json:"name,omitempty" bson:"name, omitempty"`
	Album  string             `json:"album,omitempty" bson:"album, omitempty"`
	Artist string             `json:"artist,omitempty" bson:"artist, omitempty"`
	Genres []string           `json:"genres,omitempty" bson:"genres, omitempty"`
}


