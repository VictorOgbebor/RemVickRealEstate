package main

import (
	"net/http"
	"remvick/config"
	"remvick/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSu)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/services", handlers.Repo.Services)
	mux.Get("/buisness", handlers.Repo.Buisness)
	return mux
}