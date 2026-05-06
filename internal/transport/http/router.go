package http

import (
	"github.com/KenPrz/pos-backend/internal/transport/http/util"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", util.HealthHandler)

	return r
}
