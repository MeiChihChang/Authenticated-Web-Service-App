package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.enableCORS)

	mux.Get("/", app.Home)

	mux.Post("/login", app.login)
	//mux.Get("/refresh", app.refreshToken)
	mux.Get("/logout", app.logout)

	mux.Route("/swissdata/", func(mux chi.Router){
		mux.Use(app.authRequired)

		mux.Get("/organizations", app.organization_list)
		mux.Get("/datalist/{name}", app.data_list)	
	})

	

	return mux
}