package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/s-beats/go-cms/di"
	"github.com/s-beats/go-cms/infra"
	"github.com/s-beats/go-cms/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	db, err := infra.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	di.InitRegistory(db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
