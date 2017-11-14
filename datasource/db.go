package datasource

import (
	"gopkg.in/mgo.v2"
	"log"
)

var db *mgo.Database

func GetDatabase(url string, database string) {
	session,err := mgo.Dial(url)
	if err!= nil {
		log.Fatal("db: something went wrong getting session")
	}

	db = session.DB(database)
	log.Print("db: connected to "+ db.Name)
}