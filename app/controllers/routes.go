package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
)

func (server *Server) IntializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/",Home).Methods(("GET"))

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/",http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}