//go:generate gogeneratetest -table=Verlag -endpoint=/daten/verlag/
package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	gj "github.com/quiteawful/Gjallarhorn"
)

type VerlagHandler struct {
	render         *Renderer
	verlagProvider gj.VerlagService
}

func NewVerlagHandler(v gj.VerlagService, r *Renderer) {
	return &VerlagHandler{
		render:         r,
		verlagProvider: v,
	}
}

func (h *VerlagHandler) Index(w http.ResponseWriter, r *http.Request) {
	v, err := h.verlagProvider.GetAll()
	if err != nil {
		log.Printf("could not retreive all verlag: %v\n", err)
		return
	}

	data := struct {
		Verlag []*gj.Verlag
	}{
		Verlag: v,
	}

	err = h.render.Render("base", "verlag_index", w, &data)
	if err != nil {
		log.Printf("could not execute template: %v\n", err)
		return
	}
}

func (h *VerlagHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := h.render.Render("base", "verlag_create", w, nil)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

func (h *VerlagHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	var v gj.Verlag
	var err error

	// assign values from form
	panic("scan is not implemented")

	err = h.verlagProvider.Create(&v)
	if err != nil {
		log.Printf("could not create new verlag: %v\n", err)
		return
	}

	http.Redirect(w, r, "/daten/verlag/", 301)
}

func (h *VerlagHandler) Show(w http.ResponseWriter, r *http.Request) {
	v, err := l.verlagProvider.Get(id)
	if err != nil {
		log.Printf("error while getting verlag: %v\n", err)
		return
	}

	// maybe get other providers

	data := struct {
		Verlag *gj.Verlag
	}{
		Verlag: v,
	}

	err = h.render.Render("base", "verlag_show", w, &data)
	if err != nil {
		log.Printf("error while parsing template")
		return
	}
}

func (h *VerlagHandler) Delete(w http.ResponseWriter, id int) {
	v, err := h.verlagProvider.Get(id)
	if err != nil {
		log.Printf("error while deleting verlag: %v\n", err)
		return
	}

	data := struct {
		Verlag *gj.Verlag
	}{
		Verlag: v,
	}

	err = h.render.Render("base", "verlag_delete", w, &data)
	if err != nil {
		log.Printf("error while parsing template")
		return
	}
}

func (h *VerlagHandler) DeletePOST(w http.ResponseWriter, r *http.Request) {
	_id := strings.Replace(r.URL.Path, "/daten/verlag/delete/", "", 1)
	id, err := strconv.Atoi(_id)
	if err != nil {
		log.Printf("could not parse id (/daten/verlag/%s): %v\n", _id, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Printf("could not parse delete form: %v\n", err)
		return
	}

	ok := r.FormValue("delete")
	if ok != "ok" {
		log.Printf("form value 'delete' is not 'ok', is %s\n", ok)
		return
	}

	err = h.verlag.Provider.Delete(id)
	if err != nil {
		log.Printf("could not delete verlag from db: %v\n", err)
		return
	}

	http.Redirect(w, r, "/daten/verlag/", 300)
}

func (h *VerlagHandler) Edit(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *VerlagHandler) EditPOST(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}

func (h *VerlagHandler) Search(w http.ResponseWriter, r *http.Request) {
	panic("not implemented")
}
