package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"templates-practice/pkg/config"
	"templates-practice/pkg/handlers"
	"templates-practice/pkg/render"
	"time"
)

const port = "127.0.0.1:8000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change InProduction
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create a template from cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)
	render.NewTemplate(&app)

	//http.HandleFunc("/", handlers.Repo.Homepage)

	fmt.Printf("Server started listning on %s\n", port)

	//err = http.ListenAndServe(port, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	srv := &http.Server{
		Addr:    port,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
