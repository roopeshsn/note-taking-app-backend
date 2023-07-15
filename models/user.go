package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"unique"`
	Password string
	Notes []Note `gorm:"foreignKey:Uid"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = 0
	return nil
}