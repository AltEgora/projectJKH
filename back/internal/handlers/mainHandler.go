package handlers

import (
	"ConsultantBack/internal/db/domain"
	"ConsultantBack/internal/news"
	"fmt"
	"html/template"
	"net/http"
)

type data struct {
	News []domain.New
}

func MainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on main page")

	tmpl, err := template.ParseFiles("./static/mainPage/index.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	d := data{
		News: news.UpdateNews(),
	}

	err = tmpl.Execute(w, d)
	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
