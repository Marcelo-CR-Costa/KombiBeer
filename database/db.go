package database

import (
	"log"

	"github.com/marcelocosta/kombibeer/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados(host string) {
	stringDeConexao := "host=" + host + " user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Cerveja{
		Model:       gorm.Model{},
		Nome:        "",
		Ingrediente: "",
		Preco:       "",
	})
}
