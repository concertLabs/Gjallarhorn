package web

import (
	"fmt"
	"log"
	"net/http"

	gj "github.com/quiteawful/Gjallarhorn"
)

type PersonHandler struct {
	render         *Renderer
	personProvider gj.PersonService
}

func NewPersonHandler(personProvider gj.PersonService, _render *Renderer) *PersonHandler {
	return &PersonHandler{
		personProvider: personProvider,
		render:         _render,
	}
}

func (h *PersonHandler) Index(w http.ResponseWriter, r *http.Request) {
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

func (h *PersonHandler) CreateGET(w http.ResponseWriter, r *http.Request) {
	t, err := h.render.LoadTemplate("base", "person_create")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	t.Execute(w, nil)
}

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
