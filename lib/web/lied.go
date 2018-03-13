package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/quiteawful/Gjallarhorn/lib/db"
)

type LiedHandler struct {
	render *Renderer
	db     *gorm.DB
}

func NewLiedHandler(_db *gorm.DB, r *Renderer) *LiedHandler {
	return &LiedHandler{
		db:     _db,
		render: r,
	}
}

// IndexGET shows a list with all songs in the database
func (h *LiedHandler) Index(w http.ResponseWriter, r *http.Request) {
	// TODO: add pagination

	var l []db.Lied
	if err := h.db.Find(&l).Error; err != nil {
		log.Printf("could not get all lieder: %v\n", err)
		return
	}

	data := struct {
		Lied []db.Lied
	}{
		Lied: l,
	}

	err := h.render.Render("lied_index", "lied", w, &data)
	if err != nil {
		log.Printf("error while executing template: %v\n", err)
		return
	}
}

func (h *LiedHandler) Create(w http.ResponseWriter, r *http.Request) {

	var p []db.Person
	var v []db.Verlag

	if err := h.db.Order("name, vorname").Find(&p).Error; err != nil {
		log.Printf("could not find any persons: %v\n", err)
	}

	if err := h.db.Order("name").Find(&v).Error; err != nil {
		log.Printf("could not find any verläge: %v\n", err)
	}

	data := struct {
		Personen []db.Person
		Verlag   []db.Verlag
	}{
		Personen: p,
		Verlag:   v,
	}

	err := h.render.Render("lied_create", "lied", w, &data)
	if err != nil {
		log.Printf("error while executing template: %v\n", err)
		return
	}
}

func (h *LiedHandler) CreatePOST(w http.ResponseWriter, r *http.Request) {
	// TODO: support multiple texter, arrangeurs, komponisten
	// -> Böhmischer Wind
	v := r.Form
	var l db.Lied
	var err error

	l.Titel = v.Get("titel")
	l.Untertitel = v.Get("untertitel")
	l.Jahr, err = strconv.Atoi(v.Get("jahr"))
	if err != nil {
		log.Printf("could not convert jahr to int: %v\n", err)
		l.Jahr = 0
	}
	l.KomponistID, err = strconv.Atoi(v.Get("komponist"))
	if err != nil {
		log.Printf("could not convert komponistId to int: %v\n", err)
		l.KomponistID = 0
	}
	l.TexterID, err = strconv.Atoi(v.Get("texter"))
	if err != nil {
		log.Printf("could not convert textID to int: %v\n", err)
		l.TexterID = 0
	}

	l.ArrangeurID, err = strconv.Atoi(v.Get("arrangeur"))
	if err != nil {
		log.Printf("could not convert arrangeurID to int: %v\n", err)
		l.ArrangeurID = 0
	}
	l.VerlagID, err = strconv.Atoi(v.Get("verlag"))
	if err != nil {
		log.Printf("could not convert verlagID to int: %v\n", err)
		l.VerlagID = 0
	}

	if err = h.db.Create(&l).Error; err != nil {
		log.Printf("could not create lied (%v): %v\n", l, err)
		return
	}

	http.Redirect(w, r, "/lied", 301)
}

func (h *LiedHandler) Show(w http.ResponseWriter, id uint) {
	var l db.Lied
	l.ID = id

	if err := h.db.First(&l).Error; err != nil {
		log.Printf("could not find lied %d: %v\n", id, err)
		return
	}

	var komponist, texter, arrangeur db.Person
	var verlag db.Verlag

	if err := h.db.Model(&l).Related(&komponist, "KomponistID").Error; err != nil {
		log.Printf("could not find a komponist for %s: %v\n", l.Titel, err)
	}

	if err := h.db.Model(&l).Related(&texter, "TexterID").Error; err != nil {
		log.Printf("could not find a texter for %s: %v\n", l.Titel, err)
	}

	if err := h.db.Model(&l).Related(&arrangeur, "ArrangeurID").Error; err != nil {
		log.Printf("could not find a arrangeur for %s: %v\n", l.Titel, err)
	}

	if err := h.db.Model(&l).Related(&verlag, "VerlagID").Error; err != nil {
		log.Print("could not find any verlag for %s: %v\n", l.Titel, err)
	}

	data := struct {
		Lied      *db.Lied
		Komponist *db.Person
		Text      *db.Person
		Arrangeur *db.Person
		Verlag    *db.Verlag
	}{
		Lied:      &l,
		Komponist: &komponist,
		Text:      &texter,
		Arrangeur: &arrangeur,
		Verlag:    &verlag,
	}

	err := h.render.Render("lied_show", "lied", w, &data)
	if err != nil {
		log.Printf("error while parsing template")
		return
	}
}

func (h *LiedHandler) Delete(w http.ResponseWriter, id uint) {
	var l db.Lied
	l.ID = id

	if err := h.db.First(&l).Error; err != nil {
		log.Printf("could not get lied for delete page: %v\n", err)
		return
	}

	data := struct {
		Lied *db.Lied
	}{
		Lied: &l,
	}

	err := h.render.Render("lied_delete", "lied", w, &data)
	if err != nil {
		log.Printf("error while parsing template")
		return
	}
}

func (h *LiedHandler) DeletePOST(w http.ResponseWriter, r *http.Request) {
	p := strings.Replace(r.URL.Path, "/lied/delete/", "", 1)

	id, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("could not parse id(%s) as for %s: %v\n", p, r.URL.Path, err)
		return
	}

	err = r.ParseForm()
	if err != nil {
		log.Printf("could not parse delete form: %v\n", err)
		return
	}

	// form value delete=ok??
	ok := r.FormValue("delete")
	if ok != "ok" {
		log.Printf("form value delete is not 'ok' %s\n", ok)
		return
	}

	var l db.Lied
	l.ID = uint(id)
	if err = h.db.Delete(&l).Error; err != nil {
		log.Printf("could not delete lied: %v\n", err)
		return
	}

	// TODO: change http code
	http.Redirect(w, r, "/lied", 300)
}

func (h *LiedHandler) Edit(w http.ResponseWriter, id uint) {
	var l db.Lied

	if err := h.db.First(&l, id).Error; err != nil {
		log.Printf("/lied/edit: error while getting lied: %v\n", err)
		return
	}

	if l.KomponistID != 0 {
		if err := h.db.Model(&l).Related(&l.Komponist, "KomponistID").Error; err != nil {
			log.Printf("could not find a komponist for %s: %v\n", l.Titel, err)
		}
	}

	if l.TexterID != 0 {
		if err := h.db.Model(&l).Related(&l.Texter, "TexterID").Error; err != nil {
			log.Printf("could not find a texter for %s: %v\n", l.Titel, err)
		}
	}

	if l.ArrangeurID != 0 {
		if err := h.db.Model(&l).Related(&l.Arrangeur, "ArrangeurID").Error; err != nil {
			log.Printf("could not find a arrangeur for %s: %v\n", l.Titel, err)
		}
	}

	if l.VerlagID != 0 {
		if err := h.db.Model(&l).Related(&l.Verlag, "VerlagID").Error; err != nil {
			log.Print("could not find any verlag for %s: %v\n", l.Titel, err)
		}
	}

	var _p []db.Person
	var _v []db.Verlag

	if err := h.db.Find(&_p).Error; err != nil {
		log.Printf("/lied/edit/: could not get list of personeo: %v\n", err)
	}
	if err := h.db.Find(&_v).Error; err != nil {
		log.Printf("/lied/edit/: could not get list of verläge: %v\n", err)
	}

	data := struct {
		Lied     *db.Lied
		Personen []db.Person
		Verlag   []db.Verlag
	}{
		Lied:     &l,
		Personen: _p,
		Verlag:   _v,
	}

	if err := h.render.Render("lied_edit", "lied", w, &data); err != nil {
		log.Printf("could not render edit page for %d: %v\n", id, err)
		return
	}
}

func (h *LiedHandler) EditPOST(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("could not parse edit for mof lied: %v\n", err)
		return
	}

	var l db.Lied

	x, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		log.Printf("could not convert lied id to int : %v\n", err)
		return
	}

	l.ID = uint(x)
	l.Titel = r.Form.Get("titel")
	l.Untertitel = r.Form.Get("untertitel")
	l.Jahr, err = strconv.Atoi(r.Form.Get("jahr"))

	if err != nil {
		log.Printf("could not convert jahr to int: %v\n", err)
		l.Jahr = 0
	}
	l.KomponistID, err = strconv.Atoi(r.Form.Get("komponist"))
	if err != nil {
		log.Printf("could not convert komponistId to int: %v\n", err)
		l.KomponistID = 0
	}
	l.TexterID, err = strconv.Atoi(r.Form.Get("texter"))
	if err != nil {
		log.Printf("could not convert textID to int: %v\n", err)
		l.TexterID = 0
	}

	l.ArrangeurID, err = strconv.Atoi(r.Form.Get("arrangeur"))
	if err != nil {
		log.Printf("could not convert arrangeurID to int: %v\n", err)
		l.ArrangeurID = 0
	}
	l.VerlagID, err = strconv.Atoi(r.Form.Get("verlag"))
	if err != nil {
		log.Printf("could not convert verlagID to int: %v\n", err)
		l.VerlagID = 0
	}

	if err := h.db.Save(&l).Error; err != nil {
		log.Printf("could not save lied %v\n", err)
		return
	}

	http.Redirect(w, r, "/lied", 301)
}
