package models

type Usuario struct {
	Id         int    `json:"id,omitempty"`
	Usuario    string `json:"usuario,omitempty"`
	Contrasena string `json:"contrasena,omitempty"`
	Correo     string `json:"correo,omitempty"`
}

type Usuarios []Usuario
