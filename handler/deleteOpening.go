package handler

import (
	"fmt"      //utilizado para formatar strings, aqui para formatar mensagens de erro.
	"net/http" //definir os status HTTP

	"github.com/gin-gonic/gin" //Gerenci as requisições e respostas HTTP
	"github.com/itslucasmiranda/gojob/schemas"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) { //contexto da requisição
	id := ctx.Query("id") //obtêm o parametro id da URL
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}
	opening := schemas.Opening{}                         //variável para armazenar os dados da vaga
	if err := db.First(&opening, id).Error; err != nil { //GORM para buscar as vagas pelo id
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete Opening
	if err := db.Delete(&opening).Error; err != nil { //GORM para excluir a vaga do banco de dados
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}
