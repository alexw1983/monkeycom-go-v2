package handlers

import (
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/views/dnd"
	"github.com/gorilla/mux"
)

func (h *Handler) Character(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	character := h.db.GetCharacter("alex.adwilson@gmail.com", slug)
	dnd.Character(character).Render(r.Context(), w)
}

func (h *Handler) Characters(w http.ResponseWriter, r *http.Request) {

	characters := h.db.GetCharacters("alex.adwilson@gmail.com")
	dnd.Characters(characters).Render(r.Context(), w)
}
