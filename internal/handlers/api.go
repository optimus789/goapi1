package handlers

import (
	"net/http"

	"github.com/avukadin/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.Logger)
	r.Use(chimiddle.Recoverer)
	r.Use(chimiddle.StripSlashes)
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
