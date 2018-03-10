package gjallarhorn

// Gruppe is a lookup table for person, lied, repertoire, stimmen,...
type Gruppe struct {
	ID   int
	Name string
}

type GruppenService interface {
	Get(id int) (*Gruppe, error)
	GetAll() ([]*Gruppe, error)
	Create(g *Gruppe) error
	Edit(g *Gruppe) error
	Delete(id int) error
}
