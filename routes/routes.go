package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcelocosta/kombibeer/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/cervejas", controllers.ExibeTodasCervejas)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/cervejas", controllers.CriaNovaCerveja)
	r.GET("/cervejas/:id", controllers.BuscaCervejaPorID)
	r.DELETE("/cervejas/:id", controllers.DeletaCerveja)
	r.PATCH("/cervejas/:id", controllers.EditaCerveja)
	r.GET("/cervejas/nome/:nome", controllers.BuscaCervejaPorNome)
	r.Run()
}
