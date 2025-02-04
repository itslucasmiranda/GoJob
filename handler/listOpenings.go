package handler

import (
	"net/http" //define os códigos de status HTTP nas respostas

	"github.com/gin-gonic/gin"
	"github.com/itslucasmiranda/gojob/schemas" //importa o modelo da vaga de emprego e armazena no db
)

// @BasePath /api/v1

// @Summary List openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) { //representa o contexto da requisiçaõ HTTP
	openings := []schemas.Opening{} //cria um slice (array) para armazenar as vagas encontradas no banco

	if err := db.Find(&openings).Error; err != nil { //Utiliza o ORMG GORM para buscar todas as vagas no banco de dados
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
