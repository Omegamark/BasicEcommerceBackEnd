package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	configuration "reactApp/mongoClient/config"
	"reactApp/mongoClient/db"
	handlers "reactApp/mongoClient/handler"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	duration := time.Second * 15
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	mClient, err := db.InitClient(ctx)
	if err != nil {
		fmt.Println("Failed to init client", err)
	}

	defer mClient.Disconnect(context.Background())

	conf := configuration.Config{
		DBName:   "testdb", // Change to constants
		CollName: "testCollection",
		Port:     "3000",
	}

	d := db.InitDBs(mClient, conf)

	mongoService := handlers.MongoService{
		Version: "1",
		DBCon:   d,
	}

	r := mux.NewRouter().StrictSlash(true)

	r = mongoService.Routes(r)

	fmt.Println("Server listening on %s\n", conf.Port)
	log.Fatal(http.ListenAndServe(":"+conf.Port, r))
}
