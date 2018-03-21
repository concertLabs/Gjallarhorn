package db

import "github.com/jinzhu/gorm"

type Repertoire struct {
	gorm.Model
	Name   string
	Notiz  string
	Lieder []Lied `gorm:"many2many:repertoire_lied;"`
}
