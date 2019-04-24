package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name   string `gorm:"unique;not null"`
	Passwd string `gorm:"not null"`
}

func init() {
	if !DB.HasTable(&User{}) {
		if err := DB.CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}
