package web

import (
	"log"
	"net/http"
)

func (app *WebApp) LiedIndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := app.loadTemplate("base", "lied_index")
	if err != nil {
		log.Printf("error while loading template %s: %v\n", "lied_index", err)
		return
	}

	lieder := []struct {
		Title string
	}{
		{"Test1"},
		{"Test2"},
		{"Test3"},
	}

	t.Execute(w, lieder)
}
