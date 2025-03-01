package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/yerlan/bookings/internal/config"
	"github.com/yerlan/bookings/internal/handlers"
	"github.com/yerlan/bookings/internal/render"
)

const portNum = ":8080"
var app config.AppConfig
var session *scs.SessionManager
// main is main 
func main(){

	// change this to true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port %s", portNum)

	srv := &http.Server {
		Addr: portNum,
		Handler:  routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}