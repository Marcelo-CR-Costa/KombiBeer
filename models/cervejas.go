package models

import "gorm.io/gorm"

type Cerveja struct {
	gorm.Model
	Nome        string `json:"nome"`
	Ingrediente string `json:"ingrediente"`
	Preco       string `json:"preco"` // mudar para int
}
