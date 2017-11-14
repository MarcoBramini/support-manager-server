package service

import (
	"support-manager-server/model"
	"net/http"
	"support-manager-server/datasource"
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
	users[0] = model.User{Id:"test@gmail.com", Name:"Test", Surname:"Test", Password:"test", Organization:"TestOrg"}

	user,err := datasource.GetUserById(args.UserId)
	if err!=nil {
		result.Error=err.Error()
	} else {
		result.Error="user: wrong password"
		if user.Password == args.Password {
				result.UserId = user.Id
				result.Error = ""
		}
	}
	return nil
}