package main

import (
	"os"

	"github.com/restlesswhy/rest/http-parsing-x-technology/config"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/server"
	"github.com/restlesswhy/rest/http-parsing-x-technology/pkg/logger"
	"github.com/restlesswhy/rest/http-parsing-x-technology/pkg/postgres"
)

func main() {
	logger.Info("Starting server...")
	
	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		logger.Fatalf("cant get config: %v", err)
	}

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		logger.Fatalf("Postgresql init: %s", err)
	} else {
		logger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()
	
	s := server.NewServer(cfg)

	logger.Fatal(s.Run())
}