package models

import "github.com/seunexplicit/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	FloatMap  map[string]float32
	IntMap    map[string]int
	BoolMap   map[string]bool
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
