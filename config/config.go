package config

import "html/template"

type AppConfig struct {
	UserCache     bool
	TemplateCache map[string]*template.Template
}
