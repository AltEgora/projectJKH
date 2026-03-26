package server

import (
	"ConsultantBack/back/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// Our app class
type App struct {
	server *mux.Router
}

// Constructor for App
func NewApp(addr string) *http.Server {
	rout := mux.NewRouter()

	//File server. Maybe unused in future.
	//Now needs for .scc
	rout.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	//Basic handlers
	rout.HandleFunc("/main", handlers.MainHandler).Methods("GET")

	return &http.Server{
		Handler: rout,
		Addr:    addr,
	}
}

func AddHandler(server *http.Server) {

}
