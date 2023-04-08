package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/seunexplicit/bookings/internal/config"
	"github.com/seunexplicit/bookings/internal/models"
	"github.com/seunexplicit/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About is the about page handler.
func (rep *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test":      "Hello, again!",
		"remote_ip": rep.App.Session.GetString(r.Context(), "remote_ip"),
	}

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (rep *Repository) General(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

func (rep *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (rep *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "reservation.page.html", &models.TemplateData{})
}

func (rep *Repository) Major(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

func (rep *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (rep *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (rep *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
