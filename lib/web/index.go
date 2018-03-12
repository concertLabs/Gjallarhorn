package web

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type IndexHandler struct {
	render *Renderer
	db     *gorm.DB
}

func NewIndexHandler(_db *gorm.DB, _render *Renderer) *IndexHandler {
	return &IndexHandler{
		db:     _db,
		render: _render,
	}
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {
	err := h.render.Render("index_index", "index", w, nil)
	if err != nil {
		log.Printf("error while parsing template: %v\n", err)
		return
	}
}
