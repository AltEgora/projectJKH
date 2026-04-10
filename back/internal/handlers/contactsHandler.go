package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func ContactsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on contacts page")

	tmpl, err := template.ParseFiles("/app/front/contacts.html", "/app/front/base.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", data{
		Title:  "Контакты",
		Active: "contacts",
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
