package main

import (
	"log"

	"github.com/vhe74/go-htmx/server"
)

func main() {
	log.Println("Starting....")
	s, err := server.NewServer(":3000")
	if err != nil {
		log.Fatal("Couldn't create server")
	}
	err = s.Start()
	if err != nil {
		log.Fatal("Couldn't bind server")
	}
}
