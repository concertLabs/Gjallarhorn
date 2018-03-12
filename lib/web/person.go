package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/quiteawful/Gjallarhorn/lib/db"
)

// PersonHandler is responsible for dealing with persons
type PersonHandler struct {
	render *Renderer
	db     *gorm.DB
}

// NewPersonHandler creates a new Handler for the router, personProvider is the database
// acccess layer to get and insert data, _render is our main template engine
func NewPersonHandler(_db *gorm.DB, r *Renderer) *PersonHandler {
	return &PersonHandler{
		db:     _db,
		render: r,
	}
}

// Index shows a list with all persons in the database
func (h *PersonHandler) Index(w http.ResponseWriter, r *http.Request) {
	// NOTE: add a pagination
	var p []db.Person
	if err := h.db.Find(&p).Error; err != nil {
		log.Printf("could not get all users: %v\n", err)
		return
	}
	data := struct {
		Person []db.Person
	}{
		Person: p,
	}

	err := h.render.Render("person_index", "person", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

// Create loads the html template and prints it out
func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	var g []db.Gruppe
	if err := h.db.Find(&g).Error; err != nil {
		log.Printf("could not get groups: %v\n", err)
		return
	}

	data := struct {
		Gruppen []db.Gruppe
	}{
		Gruppen: g,
	}

	err := h.render.Render("person_create", "person", w, &data)
	if err != nil {
		log.Printf("error while executing template: %v\n", err)
		return
	}
}

// Create receives the form data from CreateGET and inserts the data in the database
func (h *PersonHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	var p db.Person
	var err error
	p.Name = r.Form.Get("name")
	p.Vorname = r.Form.Get("surname")
	p.Strasse = r.Form.Get("street")
	p.PLZ = r.Form.Get("zipcode")
	p.Ort = r.Form.Get("city")
	// p.Geburtstag = r.Form.Get("birth_date")
	// p.MitgliedSeit = r.Form.Get("member_since")
	p.Email = r.Form.Get("email")
	g := r.Form.Get("gruppe")

	p.Gruppe, err = strconv.Atoi(g)
	if err != nil {
		log.Printf("PersonHandler.CreatePOST: could not convert gruppe to int: %v\n", err)
		p.Gruppe = 0
	}

	if err := h.db.Create(&p).Error; err != nil {
		log.Printf("could not create new user: %v\n", err)
		return
	}
	http.Redirect(w, r, "/person", 301)
}

// Show receives the ResponseWriter and the personID and renders the
// detail page of given person.
func (h *PersonHandler) Show(w http.ResponseWriter, id uint) {
	var p db.Person

	if err := h.db.First(&p, id).Error; err != nil {
		log.Printf("could not find user %d: %v\n", id, err)
		return
	}

	var l []db.Lied
	if err := h.db.Where("komponist_id = ? or texter_id = ? or arrangeur_id = ?", p.ID, p.ID, p.ID).Find(&l).Error; err != nil {
		log.Printf("could not find any lieder for person: %v\n", err)
	}

	data := struct {
		Person *db.Person
		Lieder []db.Lied
	}{
		Person: &p,
		Lieder: l,
	}

	err := h.render.Render("person_show", "person", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

// DeleteGET show the user a conformation page wether he really wants to delete the person
func (h *PersonHandler) Delete(w http.ResponseWriter, id uint) {
	var p db.Person
	if err := h.db.First(&p, id).Error; err != nil {
		log.Printf("could not find user for delete page %d: %v\n", id, err)
		return
	}

	data := struct {
		Person *db.Person
	}{
		Person: &p,
	}

	err := h.render.Render("person_delete", "person", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

// DeletePOST receives a conformation from DeleteGET and deltes the person
func (h *PersonHandler) DeletePOST(w http.ResponseWriter, r *http.Request) {
	p := strings.Replace(r.URL.Path, "/person/delete/", "", 1)

	id, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("could not parse id(%s) as for %s: %v\n", p, r.URL.Path, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Printf("could not parse delelte form: %v\n", err)
		return
	}

	// form value delete=ok??
	ok := r.FormValue("delete")
	if ok != "ok" {
		log.Printf("form value delete is not 'ok' %s\n", ok)
		return
	}

	var user db.Person
	user.ID = uint(id)

	if err = h.db.Delete(&user).Error; err != nil {
		log.Printf("could not delete user %d: %v\n", id, err)
		return
	}
	// TODO: change http code
	http.Redirect(w, r, "/person", 300)
}
