package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) { // *gin.Engine é um ponteiro para a instância do gin.Engine que é de forma didática um servidor web do
	v1 := router.Group("/api/v1")
	{
		// show all openings, que são as vagas de emprego
		v1.GET("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Opening",
			})
		})
		v1.POST("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "POST Opening",
			})
		})
		v1.DELETE("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE Opening",
			})
		})
		v1.PUT("/opening", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT Opening",
			})
		})
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET Openings",
			})
		})
	}
}
