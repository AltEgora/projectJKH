package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AccrualsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on accruals page")

	tmpl, err := template.ParseFiles("/app/front/calc.html", "/app/front/base.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", data{
		Active: "accruals",
		Title:  "Справочник потребителя",
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
