package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              uint       `json:"id" gorm:"primaryKey"`
	FirstName       string     `json:"firstname"`
	LastName        string     `json:"lastname"`
	Email           string     `json:"email" gorm:"unique"`
	Password        string     `json:"password"`
	Bio             *string    `json:"bio"`
	ProfilePhotoURL *string    `json:"profilephotourl"`
	Roles           []Role     `json:"roles" gorm:"many2many:user_roles"`
	Interests       []Interest `json:"interests" gorm:"many2many:user_interests"`
}
