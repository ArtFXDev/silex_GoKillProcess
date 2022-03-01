package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (server *Server) init() {
	server.Router = mux.NewRouter()
}

func (server *Server) Run(addr string) {
	server.init()
	server.initRoutes()

	fmt.Println("Listening at", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Stop(addr string) {
	server.Stop(addr)
	fmt.Println("Stopping at", addr)
}
