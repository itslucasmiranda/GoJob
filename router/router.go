package router

import (
	"os" //variáveis de ambiente, como por ex a porta do serviodor

	"github.com/gin-gonic/gin" //gerencia rotas e requisições HTTP
)

func Initialize() {
	// Inicia o servidor (router) HTTP
	router := gin.Default()

	// Inicia as rotas definidas no arquivo de rotas da API
	initializeRoutes(router)

	// define a porta em que rodará o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
