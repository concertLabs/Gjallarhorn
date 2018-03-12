package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/quiteawful/Gjallarhorn/lib/db"
)

// GruppenHandler manages our main groupable stuff
type GruppenHandler struct {
	render *Renderer
	db     *gorm.DB
}

func NewGruppenHandler(_db *gorm.DB, r *Renderer) *GruppenHandler {
	return &GruppenHandler{
		db:     _db,
		render: r,
	}
}

func (h *GruppenHandler) Index(w http.ResponseWriter, r *http.Request) {
	var g []*db.Gruppe

	if err := h.db.Find(&g).Error; err != nil {
		log.Printf("error while getting all groups: %v\n", err)
		return
	}

	data := struct {
		Gruppe []*db.Gruppe
	}{
		Gruppe: g,
	}

	h.render.Render("gruppe_index", "gruppe", w, &data)
}

func (h *GruppenHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := h.render.Render("gruppe_create", "gruppe", w, nil)
	if err != nil {
		log.Printf("error while executing template: %v\n", err)
		return
	}
}

func (h *GruppenHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	var g db.Gruppe

	if err := r.ParseForm(); err != nil {
		log.Printf("could not parse form: %v\n", err)
		return
	}
	g.Name = r.Form.Get("name")

	if err := h.db.Create(&g).Error; err != nil {
		log.Printf("error while creating new group: %v\n", err)
		return
	}

	http.Redirect(w, r, "/gruppe", 301)
}

func (h *GruppenHandler) Show(w http.ResponseWriter, id uint) {
	var group db.Gruppe
	var member []db.Person

	if err := h.db.First(&group, id).Error; err != nil {
		log.Printf("error while getting group info: %v\n", err)
		return
	}

	if err := h.db.Where("gruppe = ?", id).Find(&member).Error; err != nil {
		log.Printf("could not find any member of group %d: %v\n", id, err)
	}

	data := struct {
		Gruppe     db.Gruppe
		Mitglieder []db.Person
	}{
		Gruppe:     group,
		Mitglieder: member,
	}
	err := h.render.Render("gruppe_show", "gruppe", w, &data)
	if err != nil {
		log.Printf("could not render template: %v\n", err)
	}

}

func (h *GruppenHandler) Delete(w http.ResponseWriter, id uint) {
	var g db.Gruppe
	g.ID = id
	if err := h.db.First(&g).Error; err != nil {
		log.Printf("error while getting group for delete page: %v\n", err)
		return
	}

	data := struct {
		Gruppe *db.Gruppe
	}{
		Gruppe: &g,
	}

	err := h.render.Render("gruppe_delete", "gruppe", w, &data)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

func (h *GruppenHandler) DeletePOST(w http.ResponseWriter, r *http.Request) {
	p := strings.Replace(r.URL.Path, "/gruppe/delete/", "", 1)

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
	var g db.Gruppe
	g.ID = uint(id)
	if err = h.db.Delete(&g).Error; err != nil {
		log.Printf("error while deleting group: %v\n", err)
		return
	}

	// TODO: change http code
	http.Redirect(w, r, "/person", 300)
}
