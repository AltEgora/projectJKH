package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type data struct {
	News   []domain.ShortNew
	Phones []domain.Phone
	Title  string
	Active string
}

func MainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on main page")

	tmpl, err := template.ParseFiles("/app/front/base.html", "/app/front/home.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	newsList, err := domain.PRepo.GetShortList(context.Background(), 0, 10)
	if err != nil {
		fmt.Printf("Error while getting news from bd: %s", err)
	}

	phoneList, err := domain.PhonePRepo.GetList(context.Background())
	if err != nil {
		fmt.Printf("Error while getting news from bd: %s", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", data{
		News:   newsList,
		Phones: phoneList,
		Title:  "ПРО ЖКХ Оналйн",
		Active: "home",
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
