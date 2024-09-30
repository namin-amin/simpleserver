package server

import (
	"fmt"
	"github.com/namin-amin/simpleserver/config"
	"github.com/namin-amin/simpleserver/logger"
	"net/http"
)

type RouteHandler func(http.ResponseWriter, *http.Request) error

func (fn RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, logger logger.Logger) {
	if err := fn(w, r); err != nil {
		logger.Error("error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Server struct {
	*http.ServeMux
	config            *config.Config
	logger            logger.Logger
	globalMiddlewares []MiddlewareHandler
}

func (s *Server) HandleFunc(pattern string, handler RouteHandler, middleware ...MiddlewareHandler) {

	wrappedHandler := handler
	//first add local middleware
	for i := len(middleware) - 1; i >= 0; i-- {
		wrappedHandler = middleware[i](wrappedHandler)
	}

	//then add global middleware
	for i := len(s.globalMiddlewares) - 1; i >= 0; i-- {
		wrappedHandler = s.globalMiddlewares[i](wrappedHandler)
	}

	s.ServeMux.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		handler.ServeHTTP(writer, request, s.logger)
	})
}

func (s *Server) NewGroup(groupName string) *RouterGroup {
	return &RouterGroup{
		s:      s,
		prefix: groupName,
	}
}

func (s *Server) Run() error {
	PORT := s.config.GetEnvVarWithDefault("PORT", "5005")
	s.logger.Info("starting the server", "PORT", PORT)
	return http.ListenAndServe(fmt.Sprintf(":%s", PORT), s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ServeMux.ServeHTTP(w, r)
}

func NewServer(config *config.Config, logger logger.Logger) *Server {
	return &Server{
		ServeMux:          http.NewServeMux(),
		config:            config,
		logger:            logger,
		globalMiddlewares: []MiddlewareHandler{},
	}
}
