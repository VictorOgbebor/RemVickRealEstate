package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"remvick/config"
	"remvick/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// Sets the config for the template
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData)  {
	
	var tc map[string]*template.Template

	if app.UserCache {
	// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}

}

// Create Template Cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	tempCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tempCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tempSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return tempCache, err
		}

		matchTemps, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tempCache, err
		}

		if len(matchTemps) > 0 {
			tempSet, err = tempSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tempCache, err
			}
		}
		tempCache[name] = tempSet

	}
	return tempCache, nil
}