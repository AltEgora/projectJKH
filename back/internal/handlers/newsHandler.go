package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type dataNews struct {
	News   []domain.New
	Title  string
	Active string
}

func NewsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on news page")

	tmpl, err := template.ParseFiles("/app/front/base.html", "/app/front/news.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	newsList, err := domain.PRepo.GetList(context.Background(), 0, 10)
	if err != nil {
		fmt.Printf("Error while getting news from bd: %s", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", dataNews{
		News:   newsList,
		Title:  "Новости жкх",
		Active: "news",
	})

	if err != nil {
		fmt.Printf("Error while final step: %s", err)
	}
	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
