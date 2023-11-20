package app

import (
	"github.com/heenzu/Go-Commerce/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) IntializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods(("GET"))
}