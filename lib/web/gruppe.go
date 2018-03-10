package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	gj "github.com/quiteawful/Gjallarhorn"
)

type GruppenHandler struct {
	render          *Renderer
	gruppenProvider gj.GruppenService
}

func NewGruppenHandler(gp gj.GruppenService, r *Renderer) *GruppenHandler {
	return &GruppenHandler{
		gruppenProvider: gp,
		render:          r,
	}
}

func (h *GruppenHandler) Index(w http.ResponseWriter, r *http.Request) {
	g, err := h.gruppenProvider.GetAll()
	if err != nil {
		log.Printf("error while getting all groups: %v\n", err)
		return
	}

	data := struct {
		Gruppe []*gj.Gruppe
	}{
		Gruppe: g,
	}

	h.render.Render("base", "gruppe_index", w, &data)
}

func (h *GruppenHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := h.render.Render("base", "gruppe_create", w, nil)
	if err != nil {
		log.Printf("error while executing template: %v\n", err)
		return
	}
}

func (h *GruppenHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	var g gj.Gruppe

	g.Name = r.Form.Get("name")

	err := h.gruppenProvider.Create(&g)
	if err != nil {
		log.Printf("error while creating gruppe %s: %v\n", g.Name, err)
		return
	}

	http.Redirect(w, r, "/daten/gruppe", 301)
}

func (h *GruppenHandler) Delete(w http.ResponseWriter, id int) {
	g, err := h.gruppenProvider.Get(id)
	if err != nil {
		log.Printf("could not get gruppe (%d) while deleting: %v\n", id, err)
		return
	}

	data := struct {
		Gruppe *gj.Gruppe
	}{
		Gruppe: g,
	}

	err = h.render.Render("base", "gruppe_delete", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

func (h *GruppenHandler) DeletePOST(w http.ResponseWriter, r *http.Request) {
	p := strings.Replace(r.URL.Path, "/daten/gruppe/delete/", "", 1)

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

	err = h.gruppenProvider.Delete(id)
	if err != nil {
		log.Printf("could not delete group from db: %v\n", err)
		return
	}

	// TODO: change http code
	http.Redirect(w, r, "/person", 300)
}
