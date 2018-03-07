package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	gj "github.com/quiteawful/Gjallarhorn"
)

type LiedHandler struct {
	render         *Renderer
	liedProvider   gj.LiedService
	personProvider gj.PersonService
	verlagProvider gj.VerlagService
}

func NewLiedHandler(lp gj.LiedService, pp gj.PersonService, vp gj.VerlagService, r *Renderer) *LiedHandler {
	return &LiedHandler{
		render:         r,
		liedProvider:   lp,
		personProvider: pp,
		verlagProvider: vp,
	}
}

// IndexGET shows a list with all songs in the database
func (h *LiedHandler) Index(w http.ResponseWriter, r *http.Request) {
	// TODO: add pagination
	l, err := h.liedProvider.GetAll()
	if err != nil {
		log.Printf("could not retreive all lieder: %v\n", err)
		return
	}

	t, err := h.render.LoadTemplate("base", "lied_index")
	if err != nil {
		log.Printf("could not load templeate: %v\n", err)
		return
	}

	data := struct {
		Lied []*gj.Lied
	}{
		Lied: l,
	}

	t.Execute(w, &data)
}

func (h *LiedHandler) Create(w http.ResponseWriter, r *http.Request) {
	v := r.Form
	var l gj.Lied
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
	l.TextID, err = strconv.Atoi(v.Get("text"))
	if err != nil {
		log.Printf("could not convert textID to int: %v\n", err)
		l.TextID = 0
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

	err = h.liedProvider.Create(&l)
	if err != nil {
		log.Printf("could not create new lied: %v\n", err)
		return
	}

	http.Redirect(w, r, "/lied", 301)
}

func (h *LiedHandler) Show(w http.ResponseWriter, id int) {
	l, err := h.liedProvider.Get(id)
	if err != nil {
		log.Printf("error while  getting lied: %v\n", err)
		return
	}

	t, err := h.render.LoadTemplate("base", "lied_show")
	if err != nil {
		log.Printf("error while parsing template")
		return
	}

	var komponist, text, arrangeur *gj.Person
	var verlag *gj.Verlag

	komponist, err = h.personProvider.Get(l.KomponistID)
	if err != nil && err != gj.ErrHasNoPerson {
		log.Printf("error while getting komponist: %v\n", err)
		return
	}

	text, err = h.personProvider.Get(l.TextID)
	if err != nil && err != gj.ErrHasNoPerson {
		log.Printf("error while getting texter: %v\n", err)
		return
	}

	arrangeur, err = h.personProvider.Get(l.ArrangeurID)
	if err != nil && err != gj.ErrHasNoPerson {
		log.Printf("error while getting arrangeuer: %v\n", err)
		return
	}

	verlag, err = h.verlagProvider.Get(l.VerlagID)
	if err != nil && err != gj.ErrHasNoVerlag {
		log.Printf("error while getting verlag: %v\n", err)
		return
	}

	data := struct {
		Lied      *gj.Lied
		Komponist *gj.Person
		Text      *gj.Person
		Arrangeur *gj.Person
		Verlag    *gj.Verlag
	}{
		Lied:      l,
		Komponist: komponist,
		Text:      text,
		Arrangeur: arrangeur,
		Verlag:    verlag,
	}

	t.Execute(w, &data)
}

func (h *LiedHandler) DeleteGET(w http.ResponseWriter, id int) {
	l, err := h.liedProvider.Get(id)
	if err != nil {
		log.Printf("error while getting lied: %v\n", err)
		return
	}

	t, err := h.render.LoadTemplate("base", "lied_delete")
	if err != nil {
		log.Printf("error while parsing template")
		return
	}

	data := struct {
		Lied *gj.Lied
	}{
		Lied: l,
	}

	t.Execute(w, &data)
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

	err = h.liedProvider.Delete(id)
	if err != nil {
		log.Printf("could not delete lied from db: %v\n", err)
		return
	}

	// TODO: change http code
	http.Redirect(w, r, "/lied", 300)
}
