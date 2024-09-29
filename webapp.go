package core

import (
	"github.com/namin-amin/core/config"
	"github.com/namin-amin/core/server"
)

type WebApp interface {
	Router() *server.Server
	Config() *config.Config
}

type webApp struct {
	router *server.Server
	config *config.Config
}

func (w *webApp) Router() *server.Server {
	return w.router
}

func (w *webApp) Config() *config.Config {
	return w.config
}

func NewWebApplication() WebApp {
	config := config.NewConfig()

	return &webApp{
		router: server.NewServer(config),
	}
}
