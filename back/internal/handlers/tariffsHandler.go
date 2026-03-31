package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func TariffsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on tariffs page")

	tmpl, err := template.ParseFiles("./static/tariffs/index.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.Execute(w, data{})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
