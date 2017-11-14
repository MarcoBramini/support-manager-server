package main

import (
	"support-manager-server/service"
	"support-manager-server/datasource"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
	"support-manager-server/model"
	"fmt"
	"log"
	"encoding/base64"
)

func main() {
	datasource.GetDatabase("localhost", "support-manager")

	user := model.User{Id:"test@test.it", Name:"Test", Surname:"Test", Password:"test", Organization:"TestOrg"}
	user.Password = base64.StdEncoding.EncodeToString([]byte("test"))
	datasource.CreateUser(user)

	user,err := datasource.GetUserById("test@test.it")
	if err!=nil {
		log.Println(err.Error())
	}
	fmt.Println(user.Id)

	runServer()
}

func runServer() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	login := new(service.Login)
	s.RegisterService(login, "Login")

	router := mux.NewRouter()
	router.Handle("/", s)
	http.ListenAndServe(":3000", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET"}),
		handlers.AllowedHeaders([]string{"Content-Type"}))(router))
}
