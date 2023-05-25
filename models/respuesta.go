package models

type Respuesta struct {
	Error     bool   `json:"error"`
	Mensaje   string `json:"mensaje"`
	Resultado any    `json:"resultado"`
}

type Respuestas []Respuesta