package model

import (
	"encoding/json"
	"errors"
	"bytes"
)

type User struct {
	Mail string
	Password string
	Name string
	Surname string
	Organization string
}

type UserList []User

func (user1 User) IsEqual(user2 User) bool{
	user1ToJson,_ := json.Marshal(user1)
	user2ToJson,_ := json.Marshal(user2)
	return bytes.Compare(user1ToJson,user2ToJson)==0
}

func (slice UserList) GetByMail (mail string) (User, error){
	var u User
	for _,i := range slice {
		if i.Mail==mail{
			u = i
			return u, nil
		}
	}
	return u, errors.New("user: mail not found")
}