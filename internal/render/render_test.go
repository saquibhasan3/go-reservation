package render

import (
	"net/http"
	"testing"

	"github.com/saquibhasan3/go-reservation/internal/models"
)

func TestAddDefault(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefault(&td, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}
}

func TestTemplate(t *testing.T) {
	pathToTemplates = "./../../html/template"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	r, err := getSession()

	if err != nil {
		t.Error(err)
	}
	var ww myWriter

	err = Template(&ww, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("Error during writing template to browser")
	}

	err = Template(&ww, r, "other.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("Wrong template defined")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/home", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func TestTemplates(t *testing.T) {
	Template(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../html/template"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
