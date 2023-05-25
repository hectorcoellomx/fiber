package models

type User struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	Email  string
	Status int
}
