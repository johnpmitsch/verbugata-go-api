package router

import (
	"github.com/johnpmitsch/verbugata-go-api/api"

	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc("GET", "/", api.HandleIndex)
	return r
}
