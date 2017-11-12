package main

import (
	"support-manager-server/server"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"net/http"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	login := new(server.Login)
	s.RegisterService(login, "Login")

	router := mux.NewRouter()
	router.Handle("/", s)
	http.ListenAndServe(":3000", handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"POST", "OPTIONS", "GET"}),
		handlers.AllowedHeaders([]string{"Content-Type"}))(router))
}
