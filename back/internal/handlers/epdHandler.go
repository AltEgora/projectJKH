package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type dataEpd struct {
	States []domain.State
}

func EpdHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on epd page")

	tmpl, err := template.ParseFiles("./front/EPD.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	states, err := domain.StatePRepo.GetList(context.Background(), "/consumer/epd")
	if err != nil {
		fmt.Printf("Error while bd getting: %s\n", err)
	}

	err = tmpl.Execute(w, dataEpd{
		States: states,
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
