package router

import (
	"github.com/gin-gonic/gin"                 //gerencia as rotas
	"github.com/itslucasmiranda/gojob/docs"    //importa a doc gerada pelo swagger
	"github.com/itslucasmiranda/gojob/handler" //importa os handlers responsáveis por processar as requisições
	swaggerfiles "github.com/swaggo/files"     //package que exibe a doc
	ginSwagger "github.com/swaggo/gin-swagger" //the same
)

func initializeRoutes(router *gin.Engine) {
	// configura os handlers antes das rotas serem definidas
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{ //cada endpoint está associado a um handler que contêm a logica de processamento das requisições
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

	}
	// expõe a documentaçaõ interativa do swagger no endpoint. Através dele é possivel visualizar e testar os endpoints da API de forma gráfica
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
