package main

import (
	"doremi/internal/config"
	"doremi/internal/handler"
	"doremi/internal/repository"
	"doremi/internal/server"
	"doremi/internal/service"
	"log"
	"os"
)

func main() {
	var cfgPath string
	switch len(os.Args[1:]) {
	case 1:
		cfgPath = os.Args[1]
	case 0:
		cfgPath = "./.env"
	default:
		log.Print("USAGE: go run [CONFIG_PATH]")
		return
	}

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Print(err)
		return
	}

	db, err := repository.LoadMongo(cfg.MongoDB)
	if err != nil {
		log.Print(err)
		return
	}

	r := repository.NewRepo(db)

	s := service.NewService(*r, cfg.Auth.JwtSecretKey)

	h := handler.NewHandler(*s)

	if err := server.StartServer(cfg, h.Routes()); err != nil {
		log.Print(err)
		return
	}
}
