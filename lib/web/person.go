package web

import (
	"fmt"
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
func NewPersonHandler(personProvider gj.PersonService, _render *Renderer) *PersonHandler {
	return &PersonHandler{
		personProvider: personProvider,
		render:         _render,
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

	t, err := h.render.LoadTemplate("base", "person")
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

// CreateGET simply shows the html formular to insert a new person
func (h *PersonHandler) CreateGET(w http.ResponseWriter, r *http.Request) {
	t, err := h.render.LoadTemplate("base", "person_create")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	t.Execute(w, nil)
}

// CreatePOST receives the form data from CreateGET and inserts the data in the database
func (h *PersonHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("error while parsing form: %v\n", err)
		return
	}

	var p gj.Person
	p.Name = r.FormValue("name")
	p.Surname = r.FormValue("surname")
	p.Street = r.FormValue("street")
	p.Zipcode = r.FormValue("zipcode")
	p.City = r.FormValue("city")
	p.BirthDate = r.FormValue("birth_date")
	p.MemberSince = r.FormValue("member_since")
	p.Email = r.FormValue("email")
	p.Password = r.FormValue("password")

	err = h.personProvider.Create(&p)
	if err != nil {
		log.Printf("error while creatting user: %v\n", err)
		fmt.Fprintf(w, "error while creating user")
		return
	}

	http.Redirect(w, r, "/person", 201)
}

func (h *PersonHandler) Show(w http.ResponseWriter, r *http.Request) {
	path := strings.Replace(r.URL.Path, "/person/show/", "", 1)
	id, err := strconv.Atoi(path)
	if err != nil {
		log.Printf("error while converting id: (%s) %v\n", path, err)
		fmt.Fprintf(w, "could not convert id")
		return
	}

	p, err := h.personProvider.Get(id)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
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
func (h *PersonHandler) DeleteGET(w http.ResponseWriter, r *http.Request) {
	path := strings.Replace(r.URL.Path, "/person/delete/", "", 1)
	id, err := strconv.Atoi(path)
	if err != nil {
		log.Printf("error while converting id: (%s) %v\n", path, err)
		fmt.Fprintf(w, "could not convert id")
		return
	}

	p, err := h.personProvider.Get(id)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
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
	path := strings.Replace(r.URL.Path, "/person/delete/", "", 1)
	id, err := strconv.Atoi(path)
	if err != nil {
		log.Printf("error while converting id: (%s) %v\n", path, err)
		fmt.Fprintf(w, "could not convert id")
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Printf("could not parse delelte form: %v\n", err)
		fmt.Fprintf(w, "Person konnte nicht gelöscht werden")
		return
	}

	// form value delete=ok??
	ok := r.FormValue("delete")
	if ok != "ok" {
		log.Printf("form value delete is not 'ok' %s\n", ok)
		fmt.Fprintf(w, "Person konnte nicht gelöscht werden")
		return
	}

	err = h.personProvider.Delete(id)
	if err != nil {
		log.Printf("could not delete user from db: %v\n", err)
		fmt.Fprintf(w, "Person konnte nicht gelöscht werden")
		return
	}

	// TODO: change http code
	http.Redirect(w, r, "/person", 300)
}
