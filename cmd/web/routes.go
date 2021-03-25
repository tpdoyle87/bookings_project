package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tpdoyle87/bookings/pkg/config"
	"github.com/tpdoyle87/bookings/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//Handles panics and logs them out
	mux.Use(middleware.Recoverer)
	//handles CSRF protection for all post requests
	mux.Use(NoSurf)
	//handles loading and saving a session on every request
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}