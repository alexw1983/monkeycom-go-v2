package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/config"
	"github.com/alexw1983/monkeycom-go-v2/db"
	"github.com/alexw1983/monkeycom-go-v2/handlers"
	"github.com/alexw1983/monkeycom-go-v2/services/auth"
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

	sessionsStore := auth.NewCookieStore(auth.SessionOptions{
		CookiesKey: config.Envs.SessionCookieKey,
		MaxAge:     config.Envs.SessionMaxAge,
		HttpOnly:   config.Envs.SessionHttpOnly,
		Secure:     config.Envs.SessionSecure,
	})
	authService := auth.NewAuthService(sessionsStore)

	handlers := handlers.NewHandler(db, authService)

	// Home
	router.HandleFunc("/", auth.RequireAuth(handlers.Home, authService)).Methods("GET")

	// DnD
	router.HandleFunc("/dnd/characters", auth.RequireAuth(handlers.Characters, authService)).Methods("GET")
	router.HandleFunc("/dnd/character/{slug}", auth.RequireAuth(handlers.Character, authService)).Methods("GET")

	// Auth
	router.HandleFunc("/auth/{provider}", handlers.ProviderLogin).Methods("GET")
	router.HandleFunc("/auth/{provider}/callback", handlers.ProviderCallback).Methods("GET")
	router.HandleFunc("/auth/logout/provider", nil).Methods("GET")
	router.HandleFunc("/login", handlers.Login).Methods("GET")

	// serve files in public
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	fmt.Printf("Listening on %v\n", fmt.Sprintf("%s:%d", config.Envs.PublicHost, config.Envs.Port))

	http.ListenAndServe(fmt.Sprintf(":%d", config.Envs.Port), router)
}
