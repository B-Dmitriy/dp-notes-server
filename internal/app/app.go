package app

import (
	"log"

	"webservice/internal/config"
	"webservice/internal/db"
	"webservice/internal/server"
)

func Start() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDB(cfg.User, cfg.Password, cfg.Schema)
	if err != nil {
		log.Fatal(err)
	}

	err = server.NewServer(cfg.Host, cfg.Port, cfg.Endpoint)
	if err != nil {
		log.Fatal(err)
	}
}
