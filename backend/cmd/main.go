package main

import (
	"github.com/aternity/zense/config"
	"github.com/sirupsen/logrus"
)

//	@title			Zense
//	@version		1.0
//	@description Zense API Docs	
//	@host			localhost:8080
//	@BasePath		/api/v1

// @securityDefinitions.apiKey BearerAuth
//@in header
//@name Authorization

func main() {
	cfg, err := config.New()
	if err != nil {
		logrus.Panic(err.Error())
	}

	db, err := config.NewDatabase(cfg.Database).Connection()
	if err != nil {
		logrus.Panic(err.Error())
	}

	if err := config.NewServer(config.Server{
		Host: cfg.Server.Host,
		Port: cfg.Server.Port,
		DB:   db,
		JWT:  cfg.Server.JWT,
	}).Run(); err != nil {
		logrus.Panic(err.Error())
	}
}
