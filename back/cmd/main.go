package main

import (
	"ConsultantBack/internal/db/domain"
	"ConsultantBack/internal/db/repository"
	"ConsultantBack/internal/db/tools"
	"ConsultantBack/internal/server"
	"ConsultantBack/internal/settings"
	"fmt"
	"log"
)

func main() {
	settings.LoadEnv()

	postgres, err := tools.ConnectDb()
	if err != nil {
		fmt.Printf("Cant connect to Postgres: %s", err)
	}

	domain.PRepo = repository.NewPostRepo(postgres)
	domain.StatePRepo = repository.NewStatePostRepo(postgres)

	app := server.NewApp("0.0.0.0:" + settings.AppPort)

	fmt.Println("Server started")

	log.Fatal(app.ListenAndServe())
}
