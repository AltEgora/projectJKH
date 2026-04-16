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
	Active string
	Title  string
}

func EpdHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on glossary page")

	tmpl, err := template.ParseFiles("/app/front/glossary.html", "/app/front/base.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	states, err := domain.StatePRepo.GetList(context.Background(), "/consumer/glossary")
	if err != nil {
		fmt.Printf("Error while bd getting: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base",
		dataEpd{
			States: states,
			Active: "glossary",
			Title:  "Глоссарий ЖКХ",
		})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
