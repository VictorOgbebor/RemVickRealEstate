package main

import (
	"fmt"
	"log"
	"remvick/config"
	"remvick/pkg/handlers"
	"remvick/pkg/render"
	"time"

	"net/http"

	"github.com/alexedwards/scs/v2"
)

const PortNumber = ":8080"

// entryPoint
func main() {
	var app config.AppConfig
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	



	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	render.NewTemplates(&app)

	repo := handlers.CreateRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println(fmt.Sprintf("App is on %s", PortNumber))

	srv := &http.Server{
		Addr: PortNumber,
		Handler: Routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

/*

- Static Files
	- Consisent Templates
	- Consisent Styling
- Caching
	- Cjange templates
- Middleware
	Routing and Sessions

*/
