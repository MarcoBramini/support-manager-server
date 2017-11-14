package datasource

import (
	"log"
	"fmt"
)

type Organization struct {
	Name string `bson:"name"`
}

func GetAll(){
	var orgs []Organization
	err := db.C("organizations").Find(nil).All(&orgs)
	if err!=nil {
		log.Print(err.Error())
	}

	fmt.Println(orgs)
}