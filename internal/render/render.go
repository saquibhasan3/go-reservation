package render

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/justinas/nosurf"
	"github.com/saquibhasan3/go-reservation/internal/config"
	"github.com/saquibhasan3/go-reservation/internal/models"
)

var functions = template.FuncMap{}
var app *config.AppConfig
var pathToTemplates = "./html/template"

// New template sets
func NewRenderer(a *config.AppConfig) {
	app = a
}

func AddDefault(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	var tc map[string]*template.Template
	// Get the template cache from app config
	// Create template cache
	// tc, err := CreateTemplateCache()
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get required template from cache
	t, ok := tc[tmpl]
	if !ok {
		return errors.New("Could not get template from cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefault(td, r)

	err := t.Execute(buf, td)
	if err != nil {
		log.Fatal(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all *.page.tmpl files
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// Range through all files ending *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

/*
func RenderTemplateTest(w http.ResponseWriter, r *http.Request, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./html/template/"+tmpl, "./html/template/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, r *http.Request, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if template in cache
	_, inMap := tc[t]
	if !inMap {
		// Create Template
		log.Println("Creating template cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// call template from cache
		log.Println("Using template from cache")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./html/template/%s", t),
		"./html/template/base.layout.tmpl",
	}
	// Parsing template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add template to map or cache
	tc[t] = tmpl
	return nil
}
*/
