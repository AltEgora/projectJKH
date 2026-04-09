package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AccrualsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on accruals page")

	tmpl, err := template.ParseFiles("./front/accruals.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.Execute(w, data{})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
