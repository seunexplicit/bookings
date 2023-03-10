package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/seunexplicit/bookings/pkg/config"
	"github.com/seunexplicit/bookings/pkg/models"
)

var appConfig *config.AppConfig

func AddDefaultData(tempData *models.TemplateData) *models.TemplateData {
	return tempData
}

func LoadAppConfig(c *config.AppConfig) {
	appConfig = c
}

func RenderTemplate(w http.ResponseWriter, fileName string, tempData *models.TemplateData) {
	var template map[string]*template.Template
	if appConfig.UseCache {
		template = appConfig.TemplateCache
	} else {
		var err error

		template, err = CreateTemplateCache()
		if err != nil {
			log.Fatalf("Error creating template cache: %s", err)
			return
		}
	}

	tempFile, inMap := template[fileName]
	if !inMap {
		log.Fatal("Template not found in cache")
	}

	buff := new(bytes.Buffer)
	err := tempFile.Execute(buff, AddDefaultData(tempData))
	if err != nil {
		log.Fatal("Error: ", err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	cachedTemplate := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return cachedTemplate, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return cachedTemplate, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return cachedTemplate, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return cachedTemplate, err
			}
		}

		cachedTemplate[name] = ts
	}

	return cachedTemplate, nil
}

// func RenderTemplate(w http.ResponseWriter, fileName string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + fileName)
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Fprintf(w, fmt.Sprintf("An error occurred: %s", err))
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var err error

// 	tmpl, inMap := tc[t]
// 	if !inMap {
// 		log.Println("Printing template and adding to cache...")
// 		tmpl, err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println("Error: ", err)
// 		}
// 	} else {
// 		log.Println("loading web from cache...")
// 	}

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println("Error: ", err)
// 	}
// }

// func createTemplateCache(t string) (*template.Template, error) {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return tmpl, err
// 	}

// 	tc[t] = tmpl

// 	return tmpl, nil
// }
