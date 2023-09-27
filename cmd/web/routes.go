package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/saquibhasan3/go-reservation/internal/config"
	"github.com/saquibhasan3/go-reservation/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	/*
		mux := pat.New()
		mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	*/
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/contact", handlers.Repo.Contact)
	mux.Post("/contact", handlers.Repo.PostContact)
	mux.Get("/contact-thankyou", handlers.Repo.ContactThankyou)

	mux.Post("/search-availability", handlers.Repo.SearchAvailability)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
