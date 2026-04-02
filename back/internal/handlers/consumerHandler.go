package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func ConsumerHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on consumer page")

	tmpl, err := template.ParseFiles("./static/consumer/index.html")

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	err = tmpl.Execute(w, data{})

	if err != nil {
		fmt.Printf("Error while templating: %s\n", err)
	}

	//http.ServeFile(w, req, "./static/mainPage/index.html")
}
