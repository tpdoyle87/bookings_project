package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppCongig Holds the application level configuration
type AppConfig struct {
	UseCache			bool
	TemplateCache 		map[string]*template.Template
	Production			bool
	Session				*scs.SessionManager
}