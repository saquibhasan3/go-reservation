package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

// Struct will be available globally
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
