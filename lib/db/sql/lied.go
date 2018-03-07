package sql

import (
	"database/sql"

	gj "github.com/quiteawful/Gjallarhorn"
)

const (
	selectLied = ``
	insertLied = ``
	deleteLied = ``
)

type LiedProvider struct {
	DB *sql.DB
}

func (p *LiedProvider) Get(id int) (*gj.Lied, error) {
	panic(nil)
}

func (p *LiedProvider) GetAll() ([]*gj.Lied, error) {
	panic(nil)
}

func (p *LiedProvider) Create(l *gj.Lied) error {
	panic(nil)
}

func (p *LiedProvider) Delete(id int) error {
	panic(nil)
}

func (p *LiedProvider) Search(q string) ([]*gj.Lied, error) {
	panic(nil)
}

func (p *LiedProvider) Edit(l *gj.Lied) error {
	panic(nil)
}
