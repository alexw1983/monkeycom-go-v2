package handlers

import "github.com/alexw1983/monkeycom-go-v2/db"

type Handler struct {
	db db.DB
}

func NewHandler(db db.DB) *Handler {
	return &Handler{
		db: db,
	}
}
