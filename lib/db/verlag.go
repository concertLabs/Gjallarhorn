package db

import "github.com/jinzhu/gorm"

type Verlag struct {
	gorm.Model
	Name    string
	Strasse string
	PLZ     string
	Ort     string
}
