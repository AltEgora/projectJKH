package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type dataNews struct {
	News []domain.New
}

func NewsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on news page")

	tmpl, err := template.ParseFiles("./front/news.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	newsList, err := domain.PRepo.GetList(context.Background(), 0, 10)
	if err != nil {
		fmt.Printf("Error while getting news from bd: %s", err)
	}

	err = tmpl.Execute(w, dataNews{
		News: newsList,
	})
	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
