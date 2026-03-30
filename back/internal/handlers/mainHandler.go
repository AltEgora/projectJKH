package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type data struct {
	News []domain.New
}

func MainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on main page")

	tmpl, err := template.ParseFiles("./static/mainPage/index.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	newsList, err := domain.PRepo.GetList(context.Background(), 0, 10)
	if err != nil {
		fmt.Printf("Error while getting news from bd: %s", err)
	}

	err = tmpl.Execute(w, data{
		News: newsList,
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
