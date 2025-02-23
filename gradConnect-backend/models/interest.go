package models

type Interest struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
