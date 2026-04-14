package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func CalcHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on calc page")

	tmpl, err := template.ParseFiles("/app/front/calc.html", "/app/front/base.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", data{
		Active: "calc",
		Title:  "Калькулятор ЖКХ",
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
