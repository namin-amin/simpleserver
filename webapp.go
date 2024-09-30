package core

import (
	"github.com/namin-amin/core/config"
	"github.com/namin-amin/core/logger"
	"github.com/namin-amin/core/server"
)

type WebApp interface {
	Router() *server.Server
	Config() *config.Config
	Logger() logger.Logger
}

type webApp struct {
	router *server.Server
	config *config.Config
	logger logger.Logger
}

func (w *webApp) Router() *server.Server {
	return w.router
}

func (w *webApp) Config() *config.Config {
	return w.config
}

func (w *webApp) Logger() logger.Logger {
	return w.logger
}

func NewWebApplication() WebApp {
	config := config.NewConfig()
	log := logger.New()

	return &webApp{
		router: server.NewServer(config, log),
		config: config,
		logger: log,
	}
}
