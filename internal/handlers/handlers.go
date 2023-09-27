package handlers

import (
	"fmt"
	"net/http"

	forms "github.com/saquibhasan3/go-reservation/internal/Forms"
	"github.com/saquibhasan3/go-reservation/internal/config"
	"github.com/saquibhasan3/go-reservation/internal/driver"
	"github.com/saquibhasan3/go-reservation/internal/helpers"
	"github.com/saquibhasan3/go-reservation/internal/models"
	"github.com/saquibhasan3/go-reservation/internal/render"
	"github.com/saquibhasan3/go-reservation/internal/repository"
	"github.com/saquibhasan3/go-reservation/internal/repository/dbrepo"
)

// Repository used by handlers
var Repo *Repository

// Type of repository
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Creates new Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// Sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	first_name := r.Form.Get("first_name")
	last_name := r.Form.Get("last_name")
	email := r.Form.Get("email")
	phone := r.Form.Get("phone")
	start_date := r.Form.Get("start_date")
	end_date := r.Form.Get("end_date")
	room_id := r.Form.Get("room_id")
	adults := r.Form.Get("adults")
	children := r.Form.Get("children")
	w.Write([]byte(fmt.Sprintf("You have posted First Name : %s, Last Name : %s, Email : %s, Phone : %s, Start : %s, End : %s, Adults : %s, Children : %s, Room : %s", first_name, last_name, email, phone, start_date, end_date, adults, children, room_id)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	data := make(map[string]interface{})
	data["contact"] = contact
	render.Template(w, r, "contact.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func (m *Repository) PostContact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	contact := models.Contact{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
		Message:   r.Form.Get("message"),
	}

	form := forms.New(r.PostForm)

	// form.Has("first_name", r)
	form.Required("first_name", "last_name", "email", "phone", "message")
	form.MinLength("first_name", 5)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["contact"] = contact

		render.Template(w, r, "contact.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "contact", contact)

	http.Redirect(w, r, "/contact-thankyou", http.StatusSeeOther)

}

func (m *Repository) ContactThankyou(w http.ResponseWriter, r *http.Request) {
	contact, ok := m.App.Session.Get(r.Context(), "contact").(models.Contact)
	if !ok {
		m.App.ErrorLog.Println("Could not get contact parameters from session right now. Please try filling contact form again!")
		m.App.Session.Put(r.Context(), "error", "Could not get contact parameters from session right now. Please try filling contact form again!")
		http.Redirect(w, r, "/contact", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "contact")
	data := make(map[string]interface{})
	data["contact"] = contact
	render.Template(w, r, "thankyou.page.tmpl", &models.TemplateData{
		Data: data,
	})
}
