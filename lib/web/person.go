package web

import (
	"log"
	"net/http"

	"github.com/quiteawful/Gjallarhorn/lib/model"
)

func (a *App) person(w http.ResponseWriter, r *http.Request) {
	// /person

	// doof, aber soll uns alle geben
	p, err := model.GetPerson()
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	t, err := a.loadTemplate("base", "person")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}

	t.Execute(w, p)
}

func (a *App) personAdd(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	surname := r.FormValue("surname")

	_, err := model.NewPerson(name, surname)
	if err != nil {
		log.Printf("error while inserting new person: %v\n", err)
		return
	}

}
