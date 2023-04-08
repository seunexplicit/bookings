package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/seunexplicit/bookings/pkg"
)

type Form struct {
	url.Values
	Errors errors
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(
				field,
				fieldName(field)+" is required",
			)
		}
	}
}

func (f *Form) MinLength(field string, minLength int) {
	if len(f.Get(field)) < minLength {
		f.Errors.Add(
			field,
			fieldName(field)+" minimum length is "+fmt.Sprint(minLength),
		)
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, field+" cannot be empty")
		return false
	}

	return true
}

func fieldName(field string) string {
	return pkg.Capitalize(strings.Join(
		strings.Split(field, "_"), " ",
	))
}
