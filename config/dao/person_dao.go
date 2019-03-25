package dao

import (
	"log"

	. "github.com/marcelozilio/golang-rest-api/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PersonDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "person"
)

func (m *PersonDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *PersonDAO) GetAll() ([]Person, error) {
	var persons []Person
	err := db.C(COLLECTION).Find(bson.M{}).All(&persons)
	return persons, err
}

func (m *PersonDAO) GetByID(id string) (Person, error) {
	var person Person
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&person)
	return person, err
}

func (m *PersonDAO) Create(person Person) error {
	err := db.C(COLLECTION).Insert(&person)
	return err
}

func (m *PersonDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *PersonDAO) Update(id string, person Person) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &person)
	return err
}