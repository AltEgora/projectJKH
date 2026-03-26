package main

import (
	"ConsultantBack/back/internal/server"
	"log"
)

func main() {
	app := server.NewApp("127.0.0.1:8081")

	log.Fatal(app.ListenAndServe())
}
