package handlers

import (
	"net/http"
	"remvick/config"
	"remvick/pkg/models"
	"remvick/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates the repo for the handlers
func CreateRepo(a *config.AppConfig) *Repository  {
	return &Repository{
		App: a,
	}
}

// Sets the repo for the handlers
func NewHandlers(r *Repository)  {
	Repo = r
}

// Home is the home page handler
func (m Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

// Service is the service page handler
func (m Repository) Services(w http.ResponseWriter, r *http.Request) {
	// Perform logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"
	// send data to the template
	render.RenderTemplate(w, "services.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m Repository) Buisness(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "buisness.page.tmpl", &models.TemplateData{})
}
