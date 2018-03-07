package gjallarhorn

import "errors"

var (
	ErrHasNoVerlag = errors.New("lied has no verlag")
)

type Verlag struct {
	ID      int
	Name    string
	Street  string
	Zipcode string
	City    string
}

type VerlagService interface {
	Get(id int) (*Verlag, error)
	GetAll() ([]*Verlag, error)
	Create(v *Verlag) error
	Edit(v *Verlag) error
	Delete(id int) error
	// Search(q string) ([]*Verlag, error)
}
