package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itslucasmiranda/gojob/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) { //envia a resposta de erro para o cliente. recebe três parametros, o ctx: o context do gin, code: o código de status e msg: a mensagem enviada.
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) { //envia a resposta de sucesso. recebe três parametros, o ctx: contexto gin, o op: nome da operaçaõ realizada e data: os dados retornados na resposta.
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

// formata a mensagem de erro estruturando um padrão a ser seguido
type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

// formata a resposta de sucesso da criação
type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

// formata resposta de sucesso para exclusão
type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

// formata resposta ao buscar detalhes especificos da vaga
type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

// formata resposta ao listar todas as vagas
type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

// formata resposta após a atualização da vaga
type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
