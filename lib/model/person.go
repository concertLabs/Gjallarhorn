package model

// Person can be used as texter, musician or komponist
type Person struct {
	ID      int
	Name    string
	Surname string
}

// NewPerson creates a new person and returns the ID of this person
func NewPerson(name, surname string) (*Person, error) {
	// TODO: insert in db and return error
	// query := "insert into person(id, name, surname) values(null, ?,?)"
	return &Person{
		Name:    name,
		Surname: surname,
	}, nil
}

// GetPerson returns all Persons in the database
// TODO: add limit and offset
func GetPerson() ([]Person, error) {
	return []Person{
		Person{1, "Peter", "Pan"},
		Person{2, "Klaus-Dieter", "Wiegand"},
		Person{3, "Thomas", "Gottschalk"},
		Person{4, "Michael", "Schumacher"},
	}, nil
}
