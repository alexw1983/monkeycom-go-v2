package handlers

import (
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/views"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}
