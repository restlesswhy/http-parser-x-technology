package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/restlesswhy/rest/http-parsing-x-technology/config"
	httpdel "github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/delivery/http_del"
	postgresrepo "github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/repository/postgres_repo"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/usecase"
	"github.com/restlesswhy/rest/http-parsing-x-technology/pkg/logger"
)

const (
	maxHeaderBytes = 1 << 20
)

type Server struct {
	echo        *echo.Echo
	db          *sqlx.DB
	cfg *config.Config
	
}

func NewServer(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{
		echo: echo.New(),
		cfg: cfg,
		db: db,
	}
}

func (s *Server) Run() error {
	httpserver := &http.Server{
		Addr:           s.cfg.ServerHttp.Port,
		ReadTimeout:    time.Second * s.cfg.ServerHttp.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.ServerHttp.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	
	ctx := context.Background()
	parserRepo := postgresrepo.NewParserRepository(s.db)
	parserUC := usecase.NewParserUseCase(s.cfg, parserRepo)
	parserHandler := httpdel.NewParseHandler(parserUC)
	
	parserUC.StartParse()

	if err := s.MapHandlers(s.echo, parserHandler); err != nil {
		return err
	}

	go func() {
		logger.Infof("HTTP server is listening on PORT: %s", s.cfg.ServerHttp.Port)
		if err := s.echo.StartServer(httpserver); err != nil {
			logger.Fatal("Error starting Server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	select {
	case v := <-quit:
		logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		logger.Errorf("ctx.Done: %v", done)
	}

	if err := s.echo.Shutdown(ctx); err != nil {
		logger.Errorf("router.Shutdown: %v", err)
	}

	logger.Info("Server Exited Properly")
	return nil
}