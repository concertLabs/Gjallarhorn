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
		Person{ID: 1, Name: "Peter", Surname: "Pan"},
		Person{ID: 2, Name: "Klaus-Dieter", Surname: "Wiegand"},
		Person{ID: 3, Name: "Thomas", Surname: "Gottschalk"},
		Person{ID: 4, Name: "Michael", Surname: "Schumacher"},
	}, nil
}
