package model

// Person can be used as texter, musician or komponist
type Person struct {
	ID      int
	Name    string 
	Surname string 
}

func NewPerson(name, surname string) *Person {
	p := &Person {
		Name: name,
		Surname: surname,
	}

	return p
}

func FindPersonByID(id int) (*Person, error) {
	return nil, nil
}