package models

/*import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id     string `gorm:"primaryKey"`
	Name   string
	Email  string
	Status int
	Posts  []Post // Relación uno a muchos con la tabla Post
}*/

type User struct {
	Id     string `gorm:"primaryKey"`
	Name   string
	Email  string
	Status int
}
