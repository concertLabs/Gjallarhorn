package web

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func parseForm(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("error while parsing form %s: %v\n", r.URL.Path, err)
			return
		}

		fn(w, r)
	}
}

func plainTemplate(name string, render *Renderer) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := render.LoadTemplate("base", name)
		if err != nil {
			log.Printf("Could not parse template %s: %v\n", name, err)
			return
		}

		t.Execute(w, nil)
	}
}

func parseID(fn func(http.ResponseWriter, int), path string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p := strings.Replace(r.URL.Path, path, "", 1)

		id, err := strconv.Atoi(p)
		if err != nil {
			log.Printf("could not parse id(%s) as for %s: %v\n", p, r.URL.Path, err)
			return
		}

		fn(w, id)
	}
}
