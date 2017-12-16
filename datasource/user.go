package datasource

import (
	"support-manager-server/model"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

// CreateUser persists a new [User] into database
func CreateUser(user model.User) error {
	return db.C("users").Insert(user)
}

// GetUserById retrieves a
func GetUserById(id string) (model.User,error) {
	var user model.User

	err := db.C("users").Find(bson.M{"_id":id}).One(&user)
	if err != nil {
		return user, errors.New("db: user not found")
	}

	return user, nil
}
