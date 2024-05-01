package main

import (
	"fmt"

	"github.com/IsaelVVI/goremovebg.git/config"
	"github.com/IsaelVVI/goremovebg.git/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {

	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env")
		return
	}
	logger = config.GetLogger("main")

	logger.Debugf("Initialize Application")

	// initialize app
	router.Initialize()

}
