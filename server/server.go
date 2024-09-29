package server

import (
	"fmt"
	"net/http"

	"github.com/namin-amin/core/config"
)

type RouteHandler func(http.ResponseWriter , *http.Request) error

func (fn RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := fn(w, r); err != nil {
		fmt.Println("error: "+ err.Error())
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

type Server struct {
	*http.ServeMux
	config *config.Config
	globalMiddlewares []MiddlewareHandler
}

func  (rg *Server) HandleFunc(pattern string, handler RouteHandler, middleware ...MiddlewareHandler) ()  {

	wrappedHandler:= handler
	//first add local middleware
	for  i:=len(middleware)-1;i>=0;i--  {
		wrappedHandler = middleware[i](wrappedHandler)
	}

	//then add global middleware
	for  i:=len(rg.globalMiddlewares)-1;i>=0;i--  {
		wrappedHandler = rg.globalMiddlewares[i](wrappedHandler)
	}

	rg.ServeMux.HandleFunc(pattern, wrappedHandler.ServeHTTP)
}

func (s *Server) NewGroup(groupName string) *RouterGroup {
	return &RouterGroup{
		s:      s,
		prefix: groupName,
	}
}

func (s *Server) Run() error {
	PORT := s.config.GetEnvVarWithDefault("PORT", "5005")
	fmt.Printf("starting the server in PORT %s \n", PORT)
	return http.ListenAndServe(fmt.Sprintf(":%s", PORT), s)
}

func (s *Server)  ServeHTTP(w http.ResponseWriter,r *http.Request) ()  {
	s.ServeMux.ServeHTTP(w,r)
}

func NewServer(config *config.Config) *Server {
	return &Server{
		ServeMux: http.NewServeMux(),
		config:   config,
		globalMiddlewares: []MiddlewareHandler{},
	}
}
