package db

import "github.com/jinzhu/gorm"

type Lied struct {
	gorm.Model
	Titel      string
	Untertitel string
	Jahr       int
	// TODO: support more than one komponist
	Komponist   Person `gorm:"foreignkey:KomponistID"`
	KomponistID int
	Texter      Person `gorm:"foreignkey:TexterID"`
	TexterID    int
	Arrangeur   Person `gorm:"foreignkey:ArrangeurID"`
	ArrangeurID int
	Verlag      Verlag
	VerlagID    int
}
