package db

import "github.com/jinzhu/gorm"

func Setup(db *gorm.DB) {
	db.SingularTable(true)

	// watch for the correct orderjjj
	db.CreateTable(&Gruppe{})
	db.CreateTable(&Person{})
	db.CreateTable(&Verlag{})
	db.CreateTable(&Lied{})
}
