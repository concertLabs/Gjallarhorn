package web

import (
	"fmt"
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

	data := struct {
		Person []model.Person
	}{
		Person: p,
	}

	fmt.Printf("Data: %+v\n", data)
	t.Execute(w, &data)
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
