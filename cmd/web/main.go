package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tpdoyle87/bookings/pkg/config"
	"github.com/tpdoyle87/bookings/pkg/handlers"
	"github.com/tpdoyle87/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

var port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// Main application function and router
func main() {

	app.Production = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Name = "session_id"
	session.IdleTimeout = 20 * time.Minute
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Production

	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %v", port))
	serve := &http.Server{
		Addr: port,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
