package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//inicializa o router utilizando as configurações default do gin
	router := gin.Default() // "r" é uma variável que chamamos de efêmera, ela só  existe dentro da main e recebe o valor com :=
	//Definindo uma rota com verbos, rota "/ping" com o verbo GET e passamos uma função para essa rota
	router.GET("/ping", func(c *gin.Context) { //o handler é uma função que recebe um contexto, invoca JSON
		c.JSON(200, gin.H{ //método que retorna um JSON com status 200 e uma resposta pra quem chamou a API
			"message": "pong",
		})
	})
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
