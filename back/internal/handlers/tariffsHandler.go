package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type dataTariffs struct {
	Title  string
	Active string
}

func TariffsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on tariffs page")

	tmpl, err := template.ParseFiles("/app/front/tariffs.html", "/app/front/base.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", dataTariffs{
		Title:  "Актуальные тарифы",
		Active: "tariffs",
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
