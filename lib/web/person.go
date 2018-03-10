package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	gj "github.com/quiteawful/Gjallarhorn"
)

// PersonHandler is responsible for dealing with persons
type PersonHandler struct {
	render         *Renderer
	personProvider gj.PersonService
}

// NewPersonHandler creates a new Handler for the router, personProvider is the database
// acccess layer to get and insert data, _render is our main template engine
func NewPersonHandler(pp gj.PersonService, r *Renderer) *PersonHandler {
	return &PersonHandler{
		personProvider: pp,
		render:         r,
	}
}

// Index shows a list with all persons in the database
func (h *PersonHandler) Index(w http.ResponseWriter, r *http.Request) {
	// NOTE: add a pagination
	p, err := h.personProvider.GetAll()
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	data := struct {
		Person []*gj.Person
	}{
		Person: p,
	}

	err = h.render.Render("base", "person_index", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

// Create loads the html template and prints it out
func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := h.render.Render("base", "person_create", w, nil)
	if err != nil {
		log.Printf("error while executing template: %v\n", err)
		return
	}
}

// Create receives the form data from CreateGET and inserts the data in the database
func (h *PersonHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	var p gj.Person
	p.Name = r.Form.Get("name")
	p.Surname = r.Form.Get("surname")
	p.Street = r.Form.Get("street")
	p.Zipcode = r.Form.Get("zipcode")
	p.City = r.Form.Get("city")
	p.BirthDate = r.Form.Get("birth_date")
	p.MemberSince = r.Form.Get("member_since")
	p.Email = r.Form.Get("email")
	p.Password = r.Form.Get("password")

	err := h.personProvider.Create(&p)
	if err != nil {
		log.Printf("error while creatting user: %v\n", err)
		return
	}

	http.Redirect(w, r, "/person", 301)
}

// Show receives the ResponseWriter and the personID and renders the
// detail page of given person.
func (h *PersonHandler) Show(w http.ResponseWriter, id int) {
	p, err := h.personProvider.Get(id)
	if err != nil {
		log.Printf("error while getting person: %v\n", err)
		return
	}

	data := struct {
		Person *gj.Person
	}{
		Person: p,
	}

	err = h.render.Render("base", "person_show", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

// DeleteGET show the user a conformation page wether he really wants to delete the person
func (h *PersonHandler) Delete(w http.ResponseWriter, id int) {
	p, err := h.personProvider.Get(id)
	if err != nil {
		log.Printf("error while getting person: %v\n", err)
		return
	}

	data := struct {
		Person *gj.Person
	}{
		Person: p,
	}

	err = h.render.Render("base", "person_delete", w, &data)
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

	err = h.personProvider.Delete(id)
	if err != nil {
		log.Printf("could not delete user from db: %v\n", err)
		return
	}

	// TODO: change http code
	http.Redirect(w, r, "/person", 300)
}
