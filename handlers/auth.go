package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/views"
	"github.com/alexw1983/monkeycom-go-v2/views/auth"
	"github.com/markbates/goth/gothic"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)
	auth.Login(u).Render(r.Context(), w)
}

func (h *Handler) ProviderLogin(w http.ResponseWriter, r *http.Request) {
	if u, err := gothic.CompleteUserAuth(w, r); err == nil {
		log.Printf("User is logged in: %v", u)
		views.Home(u).Render(r.Context(), w)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
}

func (h *Handler) ProviderCallback(w http.ResponseWriter, r *http.Request) {
	u, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	log.Println(u.Email)

	err = h.auth.StoreUserSession(w, r, u)
	if err != nil {
		log.Println("Error storing user session:", err)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
