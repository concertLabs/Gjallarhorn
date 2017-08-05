package web

import (
	"fmt"
	"log"
	"net/http"
)

func (a *App) indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := a.loadTemplate("base", "index")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		fmt.Fprintf(w, "error while parsing template")
		return
	}

	t.Execute(w, nil)
}
