package server

import (
	"ConsultantBack/internal/handlers"
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
	rout.HandleFunc("/", handlers.MainHandler).Methods("GET")
	rout.HandleFunc("/tariffs", handlers.TariffsHandler).Methods("GET")
	rout.HandleFunc("/faq", handlers.FaqHandler).Methods("GET")
	rout.HandleFunc("/complaints", handlers.ComplaintsHandler).Methods("GET")
	rout.HandleFunc("/accruals", handlers.AccrualsHandler).Methods("GET")
	rout.HandleFunc("/contacts", handlers.ContactsHandler).Methods("GET")
	rout.HandleFunc("/news", handlers.NewsHandler).Methods("GET")
	rout.HandleFunc("/consumer", handlers.ConsumerHandler).Methods("GET")
	rout.HandleFunc("/consumer/epd", handlers.EpdHandler).Methods("GET")
	rout.HandleFunc("/search", handlers.SearchHandler).Methods("GET")
	rout.HandleFunc("/consumer/glossary", handlers.GlossaryHandler).Methods("GET")
	rout.HandleFunc("/consumer/jku", handlers.JkuHandler).Methods("GET")
	rout.HandleFunc("/consumer/serviceproviders", handlers.ServiceprovidersHandler).Methods("GET")
	//rout.HandleFunc("/calc", handlers.CalcHandler).Methods("GET")

	rout.HandleFunc("/api/gis", handlers.GisHandler).Methods("POST")

	return &http.Server{
		Handler: rout,
		Addr:    addr,
	}
}

func AddHandler(server *http.Server) {

}
