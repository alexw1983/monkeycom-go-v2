package handlers

import (
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/views/dnd"
	"github.com/gorilla/mux"
)

func (h *Handler) Character(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)

	vars := mux.Vars(r)
	slug := vars["slug"]
	character := h.db.GetCharacter(u.Email, slug)
	dnd.Character(u, character).Render(r.Context(), w)
}

func (h *Handler) Characters(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)
	characters := h.db.GetCharacters(u.Email)

	dnd.Characters(u, characters).Render(r.Context(), w)
}
