package user

import (
	"rest-go/internal/infra/httpx"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r chi.Router, handler *Handler) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", httpx.Wrap(handler.Create))
		r.Get("/", httpx.Wrap(handler.List))
		r.Get("/{id}", httpx.Wrap(handler.Get))
		r.Put("/{id}", httpx.Wrap(handler.Update))
		r.Delete("/{id}", httpx.Wrap(handler.Delete))
	})
}
