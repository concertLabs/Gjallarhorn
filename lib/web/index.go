package web

import (
	"fmt"
	"log"
	"net/http"
)

type IndexHandler struct {
	render *Renderer
}

func NewIndexHandler(_render *Renderer) *IndexHandler {
	return &IndexHandler{render: _render}
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	t, err := h.render.LoadTemplate("base", "index")
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		fmt.Fprintf(w, "error while parsing template")
		return
	}

	t.Execute(w, nil)
}
