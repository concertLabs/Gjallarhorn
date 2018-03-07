package sql

import (
	"database/sql"

	gj "github.com/quiteawful/Gjallarhorn"
)

const (
	selectVerlag = ``
	insertVerlag = ``
	deleteVerlag = ``
)

type VerlagProvider struct {
	DB *sql.DB
}

func (p *VerlagProvider) Get(id int) (*gj.Verlag, error) {
	panic(nil)
}

func (p *VerlagProvider) GetAll() ([]*gj.Verlag, error) {
	panic(nil)
}

func (p *VerlagProvider) Create(v *gj.Verlag) error {
	panic(nil)
}

func (p *VerlagProvider) Delete(id int) error {
	panic(nil)
}

func (p *VerlagProvider) Search(q string) ([]*gj.Verlag, error) {
	panic(nil)
}

func (p *VerlagProvider) Edit(v *gj.Verlag) error {
	panic(nil)
}
