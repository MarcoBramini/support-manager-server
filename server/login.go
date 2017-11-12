package server

import (
	"support-manager/model"
	"net/http"
)

type LoginRequest struct {
	UserId string `json:"userId"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserId string `json:"userId"`
	Error string `json:"error"`
}

type Login string

func (t *Login) DoLogin (r *http.Request, args *LoginRequest, result *LoginResponse) error{
	var users = make(model.UserList, 1)
	users[0] = model.User{Mail:"test@gmail.com", Name:"Test", Surname:"Test", Password:"test", Organization:"TestOrg"}

	user,err := users.GetByMail(args.UserId)
	if err!=nil {
		result.Error="user: not found" + err.Error()
	} else {
		result.Error="user: wrong password"
		if user.Password == args.Password {
				result.UserId = user.Mail
				result.Error = ""
		}
	}
	return nil
}