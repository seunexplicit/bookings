package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/seunexplicit/bookings/cmd/middlewares"
	"github.com/seunexplicit/bookings/pkg/config"
	"github.com/seunexplicit/bookings/pkg/handlers"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middlewares.PrintPath)
	mux.Use(middlewares.CSRFHandler)
	mux.Use(middlewares.LoadSession)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
