package main // ponto de partida da aplicação. ao executar o go procura a main para executar

import (
	"github.com/itslucasmiranda/gojob/config" //fornece as configurações para a aplicação, como variáveis de ambiente e parametros.
	"github.com/itslucasmiranda/gojob/router" //contêm a lógica relacionada ao roteamento das requisições HTTP da API, onde configuramos EndPoints.
)

var (
	logger *config.Logger //registra mensagens de log, como erros ou mensagens
)

// incializa-se a execução do logger
func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
