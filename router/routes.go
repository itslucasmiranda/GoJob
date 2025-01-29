package router

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) { // *gin.Engine é um ponteiro para a instância do gin.Engine que é de forma didática um servidor web do
	v1 := router.Group("/api/v1")
	{
		//Aqui é onde estabelecemos as rotas vindas do arquivo handler.go que é responsável por registrar as nossas rotas
		v1.GET("/opening", handlers ... :  handler.ShowOpeningHandler)
		v1.POST("/opening", handlers ... : handler.CreateOpeningHandler)
		v1.DELETE("/opening", handlers ... : handler.DeleteOpeningHandler)
		v1.PUT("/opening", handlers ... : handler.UpdateOpeningHandler)
		v1.GET("/openings", handlers ... : handler.ListOpeningsHandler)

	}
}
