package model

import (
	"encoding/json"
	"bytes"
)

type User struct {
	Id string `bson:"_id"`
	Password string `bson:"password"`
	Name string `bson:"name"`
	Surname string `bson:"surname"`
	Organization string `bson:"organization"`
}

type UserList []User

func (user1 User) IsEqual(user2 User) bool{
	user1ToJson,_ := json.Marshal(user1)
	user2ToJson,_ := json.Marshal(user2)
	return bytes.Compare(user1ToJson,user2ToJson)==0
}