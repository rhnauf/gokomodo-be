package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/rhnauf/gokomodo-be/internal/api"
)

func run() error {
	log.Println("starting up application")

	// load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env files")
	}

	httpHandler := api.NewHandler()
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}
