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

	t, err := h.render.LoadTemplate("base", "person_index")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	data := struct {
		Person []*gj.Person
	}{
		Person: p,
	}

	t.Execute(w, &data)
}

// Create receives the form data from CreateGET and inserts the data in the database
func (h *PersonHandler) Create(w http.ResponseWriter, r *http.Request) {
	v := r.Form
	// TODO: change to v url.Values
	var p gj.Person
	p.Name = v.Get("name")
	p.Surname = v.Get("surname")
	p.Street = v.Get("street")
	p.Zipcode = v.Get("zipcode")
	p.City = v.Get("city")
	p.BirthDate = v.Get("birth_date")
	p.MemberSince = v.Get("member_since")
	p.Email = v.Get("email")
	p.Password = v.Get("password")

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

	t, err := h.render.LoadTemplate("base", "person_show")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	data := struct {
		Person *gj.Person
	}{
		Person: p,
	}

	t.Execute(w, &data)
}

// DeleteGET show the user a conformation page wether he really wants to delete the person
func (h *PersonHandler) DeleteGET(w http.ResponseWriter, id int) {
	p, err := h.personProvider.Get(id)
	if err != nil {
		log.Printf("error while getting person: %v\n", err)
		return
	}

	t, err := h.render.LoadTemplate("base", "person_delete")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	data := struct {
		Person *gj.Person
	}{
		Person: p,
	}

	t.Execute(w, &data)
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
