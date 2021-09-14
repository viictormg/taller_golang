package main

import (
	// "encoding/json"
	// "fmt"

	"example.com/taller/src/database"
	"example.com/taller/src/routes"

	_ "github.com/lib/pq"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	config := atreugo.Config{
		Addr: "localhost:9092",
	}
	server := atreugo.New(config)
	router := server.NewGroupPath("/v1")
	routes.SetRouter(router)

	database.GetConnection()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
