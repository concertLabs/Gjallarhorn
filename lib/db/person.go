package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	Name    string
	Vorname string
	Strasse string
	PLZ     string
	Ort     string

	Geburtstag   time.Time
	MitgliedSeit time.Time

	Gruppe int

	Email string
	Rolle int
}
