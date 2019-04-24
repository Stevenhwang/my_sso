package models

import "github.com/jinzhu/gorm"

type Door struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Operator string `gorm:"not null"`
	Status   int    `gorm:"default:1"`
	Img      string `gorm:"unique;not null"`
	Link     string `gorm:"not null"`
}

func init() {
	if !DB.HasTable(&Door{}) {
		if err := DB.CreateTable(&Door{}).Error; err != nil {
			panic(err)
		}
	}
}
