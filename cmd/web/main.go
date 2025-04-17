package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/config"
	"github.com/alexw1983/monkeycom-go-v2/db"
	"github.com/alexw1983/monkeycom-go-v2/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	// server := api.NewAPIServer(":8080", db)
	// if err := server.Run(); err != nil {
	// 	log.Fatal(err)
	// }

	router := mux.NewRouter()
	handlers := handlers.NewHandler(db)

	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/dnd/characters", handlers.Characters).Methods("GET")
	router.HandleFunc("/dnd/character/{slug}", handlers.Character).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", fmt.Sprintf("%s:%d", config.Envs.PublicHost, config.Envs.Port))

	http.ListenAndServe(fmt.Sprintf(":%d", config.Envs.Port), router)
}
