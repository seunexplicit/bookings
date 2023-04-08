package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/seunexplicit/bookings/cmd/middlewares"
	"github.com/seunexplicit/bookings/internal/config"
	"github.com/seunexplicit/bookings/internal/handlers"
	"github.com/seunexplicit/bookings/internal/render"
)

const port = ":8080"

var appConfig config.AppConfig

func main() {
	appConfig = loadTemplate()
	appConfig.UseCache = false
	appConfig.InProduction = false
	appConfig.Session = setSession()

	render.LoadAppConfig(&appConfig)
	middlewares.LoadAppConfig(&appConfig)

	repo := handlers.NewRepository(&appConfig)
	handlers.NewHandler(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port " + port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&appConfig),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loadTemplate() config.AppConfig {
	fileTemplates, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	return config.AppConfig{
		TemplateCache: fileTemplates,
	}
}

func setSession() *scs.SessionManager {
	scsManager := scs.New()
	scsManager.Lifetime = 24 * time.Hour
	scsManager.Cookie = scs.SessionCookie{
		Name:     "session",
		Persist:  true,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	return scsManager
}
