package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/maxzycon/SIB-Golang-Final-Project-3/cmd"
	"github.com/maxzycon/SIB-Golang-Final-Project-3/internal/config"
)

func main() {
	config.Init()
	conf := config.Get()
	log.Info("[InitialEnv] env set successfully")
	cmd.InitWebservice(&cmd.InitWebserviceParam{Conf: conf})
}
