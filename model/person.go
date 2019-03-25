package model

import "gopkg.in/mgo.v2/bson"

type Person struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Sex   		string        `bson:"sex" json:"sex"`
	BirthDate	string        `bson:"birth_date" json:"birth_date"`
}