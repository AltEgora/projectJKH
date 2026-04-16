package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type dataServiceproviders struct {
	States []domain.State
	Active string
	Title  string
}

func ServiceprovidersHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on serviceproviders page")

	tmpl, err := template.ParseFiles("/app/front/serviceproviders.html", "/app/front/base.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	states, err := domain.StatePRepo.GetList(context.Background(), "/consumer/serviceproviders")
	if err != nil {
		fmt.Printf("Error while bd getting: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base",
		dataServiceproviders{
			States: states,
			Active: "serviceproviders",
			Title:  "Поставщики услуг и органы контроля",
		})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
