package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/s-beats/go-cms/di"
	"github.com/s-beats/go-cms/infra/db"
	"github.com/s-beats/go-cms/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	dbClient, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	di.InitRegistory(dbClient)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
