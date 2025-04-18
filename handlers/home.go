package handlers

import (
	"net/http"

	"github.com/alexw1983/monkeycom-go-v2/views"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	u, _ := h.auth.GetUserSession(r)
	views.Home(u).Render(r.Context(), w)
}
