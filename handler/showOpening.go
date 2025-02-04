package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itslucasmiranda/gojob/schemas" //importa o modelo da vaga de emprego e armazena no db
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(ctx *gin.Context) { //represeta o contextoda requisição HTTP
	id := ctx.Query("id") //obtêm o parametro id da URL
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}                         //cria uma variável para armazenar os dados da vaga encontrada
	if err := db.First(&opening, id).Error; err != nil { //busca a vaga no banco pelo id informado
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSuccess(ctx, "show-opening", opening)
}
