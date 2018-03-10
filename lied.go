package gjallarhorn

import "errors"

var (
	ErrHasNoPerson = errors.New("lied has no person")
)

type Lied struct {
	ID         int    `sql:"id"`
	Titel      string `sql:"titel"`
	Untertitel string `sql:"untertitel"`
	Jahr       int    `sql:"jahr"`
	// TODO: make a slice with komponist, text and arrangeur
	KomponistID int `sql:"komponist_id"`
	TextID      int `sql:"text_id"`
	ArrangeurID int `sql:"arrangeur_id"`
	VerlagID    int `sql:"verlag_id"`
}

type LiedService interface {
	Get(id int) (*Lied, error)
	GetAll() ([]*Lied, error)
	Create(l *Lied) error
	Edit(l *Lied) error
	Delete(id int) error
	Search(q string) ([]*Lied, error)
}
