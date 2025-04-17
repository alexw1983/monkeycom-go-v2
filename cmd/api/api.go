package api

import (
	"log"
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/db"
	"github.com/alexw1983/monkeycom-go-v2/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   db.DB
}

func NewAPIServer(addr string, db db.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (server *APIServer) Run() error {

	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", server.addr)

	return http.ListenAndServe(server.addr, router)
}
