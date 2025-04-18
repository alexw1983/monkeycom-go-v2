package handlers

import (
	"github.com/alexw1983/monkeycom-go-v2/db"
	"github.com/alexw1983/monkeycom-go-v2/services/auth"
)

type Handler struct {
	db   db.DB
	auth *auth.AuthService
}

func NewHandler(db db.DB, auth *auth.AuthService) *Handler {
	return &Handler{
		db:   db,
		auth: auth,
	}
}
