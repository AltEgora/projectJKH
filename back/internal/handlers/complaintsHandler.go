package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func ComplaintsHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on complaints page")

	tmpl, err := template.ParseFiles("./static/complaints/index.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.Execute(w, data{})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
