package gjallarhorn

import "time"

type Person struct {
	ID      int    `sql:"id"`
	Name    string `sql:"name"`
	Surname string `sql:"surname"`
	Street  string `sql:"street"`
	Zipcode string `sql:"zipcode"`
	City    string `sql:"city"`

	// TOOD: make timestampe out of it
	BirthDate   string `sql:"birth_date"`
	MemberSince string `sql:"member_since"`

	Group int `sql:"group"`

	Email       string    `sql:"email"`
	Password    string    `sql:"password"`
	AccessLevel int       `sql:"access_level"`
	CreatedAt   time.Time `sql:"created_at"`
}

// TODO: add a group field to database and frontend

type PersonService interface {
	Get(id int) (*Person, error)
	GetAll() ([]*Person, error)
	Create(p *Person) error
	// Edit(p *Person) error
	Delete(id int) error
	Search(q string) ([]*Person, error)
}
