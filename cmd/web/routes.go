package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/seunexplicit/bookings/cmd/middlewares"
	"github.com/seunexplicit/bookings/internal/config"
	"github.com/seunexplicit/bookings/internal/handlers"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(chimiddleware.Recoverer)
	mux.Use(middlewares.PrintPath)
	mux.Use(middlewares.CSRFHandler)
	mux.Use(middlewares.LoadSession)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/general-quaters", handlers.Repo.General)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/major-suites", handlers.Repo.Major)
	mux.Get("/search-availability", handlers.Repo.Availability)

	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	staticFilePath := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", staticFilePath))

	return mux
}
