package server

import (
	"fmt"
	"net/http"

	"github.com/namin-amin/simpleserver/config"
	"github.com/namin-amin/simpleserver/logger"
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
		wrappedHandler.ServeHTTP(writer, request, s.logger)
	})
}

func (s *Server) GET(route string, 
					handler RouteHandler, 
					middleware ...MiddlewareHandler) {
	s.HandleFunc(fmt.Sprintf("GET %s", route), handler, middleware...)
}

func (s *Server) POST(route string, 
					hadler RouteHandler, 
					middleware ...MiddlewareHandler) {
	s.HandleFunc(fmt.Sprintf("POST %s", route), hadler, middleware...)
}

func (s *Server) PUT(route string, 
					hanler RouteHandler, 
					middleware ...MiddlewareHandler) {
	s.HandleFunc(fmt.Sprintf("PUT %s", route), hanler, middleware...)
}

func (s *Server) PATCH(route string, 
					hanler RouteHandler, 
					middleware ...MiddlewareHandler) {
	s.HandleFunc(fmt.Sprintf("PATCH %s", route), hanler, middleware...)
}

func (s *Server) DELETE(route string, 
						hanler RouteHandler, 
						middleware ...MiddlewareHandler) {
	s.HandleFunc(fmt.Sprintf("DELETE %s", route), hanler, middleware...)
}

func (s *Server) HEAD(route string, 
					hanler RouteHandler, 
					middleware ...MiddlewareHandler) {
	s.HandleFunc(fmt.Sprintf("HEAD %s", route), hanler, middleware...)
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
