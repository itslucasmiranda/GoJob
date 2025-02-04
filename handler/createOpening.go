package handler

import (
	"net/http" //contêm os códigos do status HTTP usados nas respostas

	"github.com/gin-gonic/gin"                 //gerencia as requisições e respostas HTTP
	"github.com/itslucasmiranda/gojob/schemas" //importa o modelo da vaga de emprego e armazena no db
)

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) { //permite acessar os dados recebidos
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request) //converte os dados JSON para a estrutura do create

	if err := request.Validate(); err != nil { //verifica se os campos foram preenchidos corretamente
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil { //insere a vaga de emprego no banco de dados usando GORM
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
