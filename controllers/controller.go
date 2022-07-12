package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcelocosta/kombibeer/database"
	"github.com/marcelocosta/kombibeer/models"
)

func ExibeTodasCervejas(c *gin.Context) {
	var cervejas []models.Cerveja
	database.DB.Find(&cervejas)
	c.JSON(200, cervejas)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "Olá " + nome + ",bem vindo a Kombibeer",
	})

}
func CriaNovaCerveja(c *gin.Context) {
	var cerveja models.Cerveja
	if err := c.ShouldBindJSON(&cerveja); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&cerveja)
	c.JSON(http.StatusOK, cerveja)
}
func BuscaCervejaPorID(c *gin.Context) {
	var cerveja models.Cerveja
	id := c.Params.ByName("id")
	database.DB.First(&cerveja, id)

	if cerveja.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto não Cadastrado"})
		return
	}

	c.JSON(http.StatusOK, cerveja)
}

func DeletaCerveja(c *gin.Context) {
	var cerveja models.Cerveja
	id := c.Params.ByName("id")
	database.DB.Delete(&cerveja, id)
	c.JSON(http.StatusOK, gin.H{"data": "Produto deletado com sucesso"})
}

func EditaCerveja(c *gin.Context) {
	var cerveja models.Cerveja
	id := c.Params.ByName("id")
	database.DB.First(&cerveja, id)

	if err := c.ShouldBindJSON(&cerveja); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&cerveja).UpdateColumns(cerveja)
	c.JSON(http.StatusOK, cerveja)
}

func BuscaCervejaPorNome(c *gin.Context) {
	var cerveja models.Cerveja
	nome := c.Param("nome")
	database.DB.Where(&models.Cerveja{Nome: nome}).First(&cerveja)

	if cerveja.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Produto não Cadastrado"})
		return
	}
	c.JSON(http.StatusOK, cerveja)

}
