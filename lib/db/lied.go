package db

import "github.com/jinzhu/gorm"

type Lied struct {
	gorm.Model
	Titel      string
	Untertitel string
	Jahr       int
	// TODO: support more than one komponist
	Komponisten []Person `gorm:"many2many:lied_person;"`
	Texter      []Person `gorm:"many2many:lied_person;"`
	Arrangeur   []Person `gorm:"many2many:lied_person;"`
	Verlag      []Verlag `gorm:"many2many:lied_verlag;"`
}
