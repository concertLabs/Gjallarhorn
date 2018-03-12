//go:generate gogeneratetest -table=Verlag -endpoint=/daten/verlag/
package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/quiteawful/Gjallarhorn/lib/db"
)

type VerlagHandler struct {
	render *Renderer
	db     *gorm.DB
}

func NewVerlagHandler(_db *gorm.DB, r *Renderer) *VerlagHandler {
	return &VerlagHandler{
		db:     _db,
		render: r,
	}
}

func (h *VerlagHandler) Index(w http.ResponseWriter, r *http.Request) {
	var v []db.Verlag

	if err := h.db.Find(&v).Error; err != nil {
		log.Printf("error while getting all verläge: %v\n", err)
		return
	}

	data := struct {
		Verlag []db.Verlag
	}{
		Verlag: v,
	}

	err := h.render.Render("verlag_index", "verlag", w, &data)
	if err != nil {
		log.Printf("could not execute template: %v\n", err)
		return
	}
}

func (h *VerlagHandler) Create(w http.ResponseWriter, r *http.Request) {
	err := h.render.Render("verlag_create", "verlag", w, nil)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}

func (h *VerlagHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	var v db.Verlag

	v.Name = r.Form.Get("name")
	v.Strasse = r.Form.Get("street")
	v.PLZ = r.Form.Get("zipcode")
	v.Ort = r.Form.Get("city")

	if err := h.db.Create(&v).Error; err != nil {
		log.Printf("error while creating new verlag: %v\n", err)
		return
	}

	http.Redirect(w, r, "/verlag", 301)
}

func (h *VerlagHandler) Show(w http.ResponseWriter, id uint) {
	var v db.Verlag
	if err := h.db.First(&v, id).Error; err != nil {
		log.Printf("error while getting verlag %d: %v\n", id, err)
		return
	}

	var l []db.Lied
	if err := h.db.Where("verlag_id = ?", v.ID).Find(&l).Error; err != nil {
		log.Printf("error while getting lieder from verlag: %v\n", err)
		// maaaay return
	}

	data := struct {
		Verlag *db.Verlag
		Lieder []db.Lied
	}{
		Verlag: &v,
		Lieder: l,
	}

	err := h.render.Render("verlag_show", "verlag", w, &data)
	if err != nil {
		log.Printf("error while parsing template")
		return
	}
}

func (h *VerlagHandler) Delete(w http.ResponseWriter, id uint) {
	var v db.Verlag
	if err := h.db.First(&v, id).Error; err != nil {
		log.Printf("error while getting verlag %d: %v\n", id, err)
		return
	}

	data := struct {
		Verlag *db.Verlag
	}{
		Verlag: &v,
	}

	err := h.render.Render("verlag_delete", "verlag", w, &data)
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

	var v db.Verlag
	v.ID = uint(id)

	if err = h.db.Delete(&v).Error; err != nil {
		log.Printf("error while deleting verlag: %v\n", err)
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
