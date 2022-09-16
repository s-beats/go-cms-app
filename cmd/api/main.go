package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/s-beats/go-cms/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
