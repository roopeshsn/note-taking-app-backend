package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Uid uint `gorm:"index"`
	Note string
}

func (n *Note) TableName() string {
	return "notes"
}