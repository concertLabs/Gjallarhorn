package web

import (
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
	err := h.render.Render("index_index", "index", w, nil)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}
