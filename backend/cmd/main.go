package main

import (
	"github.com/aternity/zense/config"
	"github.com/sirupsen/logrus"
)

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
		JWT:  cfg.Server.JWT,
		DB:   db,
	}).Run(); err != nil {
		logrus.Panic(err.Error())
	}
}
