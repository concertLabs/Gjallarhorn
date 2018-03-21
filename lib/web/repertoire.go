package web

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/quiteawful/Gjallarhorn/lib/db"
)

type RepertoireHandler struct {
	render *Renderer
	db     *gorm.DB
}

func NewRepertoireHandler(_db *gorm.DB, r *Renderer) *RepertoireHandler {
	return &RepertoireHandler{
		db:     _db,
		render: r,
	}
}

func (h *RepertoireHandler) Index(w http.ResponseWriter, r *http.Request) {
	var repertoires []db.Repertoire

}
