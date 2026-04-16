package handlers

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"fmt"
	"html/template"
	"net/http"
)

type dataJku struct {
	States []domain.State
	Active string
	Title  string
}

func JkuHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request on JKU page")

	tmpl, err := template.ParseFiles(
		"/app/front/JKU.html",
		"/app/front/base.html",
	)
	if err != nil {
		fmt.Printf("Template error: %s\n", err)
		return
	}

	states, err := domain.StatePRepo.GetList(context.Background(), "/consumer/jku")
	if err != nil {
		fmt.Printf("DB error: %s\n", err)
	}

	err = tmpl.ExecuteTemplate(w, "base", struct {
		States []domain.State
		Active string
		Title  string
	}{
		States: states,
		Active: "jku",
		Title:  "ЖКУ",
	})

	if err != nil {
		fmt.Printf("Execute error: %s\n", err)
	}
}
