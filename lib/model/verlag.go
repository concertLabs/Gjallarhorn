package model

import "errors"

// Verlag asdf
type Verlag struct {
	ID      int
	Name    string
	Zusatz  string
	Strasse string
	PLZ     string
	Ort     string
}

// NewVerlag can be used to store a new verlag in the database
func NewVerlag(name string) (*Verlag, error) {
	return &Verlag{Name: name}, nil
}

// CreateVerlag inserts the new verlag in the databse
func CreateVerlag(v Verlag) error {
	if v.Name == "" {
		return errors.New("name is empty")
	}

	// q := fmt.Sprintf("insert into verlag(name, zusatz, strasse, plz, ort) values(?, ?, ?, ?, ?);")

	return nil
}
