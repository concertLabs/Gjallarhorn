package gjallarhorn

import "errors"

var (
	ErrHasNoPerson = errors.New("lied has no person")
)

type Lied struct {
	ID         int
	Titel      string
	Untertitel string
	Jahr       int
	// TODO: make a slice with komponist, text and arrangeur
	KomponistID int
	TextID      int
	ArrangeurID int
	VerlagID    int
}

type LiedService interface {
	Get(id int) (*Lied, error)
	GetAll() ([]*Lied, error)
	Create(l *Lied) error
	Edit(l *Lied) error
	Delete(id int) error
	Search(q string) ([]*Lied, error)
}
