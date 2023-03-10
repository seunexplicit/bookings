package handlers

import (
	"net/http"

	"github.com/seunexplicit/bookings/pkg/config"
	"github.com/seunexplicit/bookings/pkg/models"
	"github.com/seunexplicit/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (rep *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	rep.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler.
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test":      "Hello, again!",
		"remote_ip": rep.App.Session.GetString(r.Context(), "remote_ip"),
	}

	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
