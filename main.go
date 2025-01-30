package main // package é como se fossem subprojetos dentro do nosso projeto

import (
	"github.com/itslucasmiranda/gojob/config"
	"github.com/itslucasmiranda/gojob/router"
)

var (
	logger *config.Logger
)

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
