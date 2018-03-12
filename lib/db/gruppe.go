package db

import "github.com/jinzhu/gorm"

type Gruppe struct {
	gorm.Model
	Name string
}
