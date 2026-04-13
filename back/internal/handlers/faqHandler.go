package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func FaqHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on faq page")

	tmpl, err := template.ParseFiles("/app/front/FAQ.html", "/app/front/base.html")
	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", data{
		Title:  "Самые частые вопросы",
		Active: "faq",
	})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
