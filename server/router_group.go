package server

import (
	"fmt"
	"net/http"
	"strings"
)

type RouterGroup struct {
	s      *Server
	prefix string
}

func (rg *RouterGroup) Handle(pattern string, handler http.Handler) {
	patternContent := strings.Split(pattern, " ")
	if len(patternContent) == 2 {
		rg.s.Handle(
			fmt.Sprintf("%s %s", 
						patternContent[0], 
						rg.prefix+patternContent[1]), 
						handler)
		return
	}

	rg.s.Handle(rg.prefix+pattern, handler)
}

func (rg *RouterGroup) HandleFunc(pattern string, handler RouteHandler, middleware ...MiddlewareHandler) {
	patternContent := strings.Split(pattern, " ")
	if len(patternContent) == 2 {
			rg.s.HandleFunc(
				fmt.Sprintf("%s %s", 
							patternContent[0], 
							rg.prefix+patternContent[1]), 
							handler,
							middleware...)
		return
	}

	rg.s.HandleFunc(rg.prefix+pattern, handler, middleware...)
}

func (s *RouterGroup) GET(route string, hanler RouteHandler, middleware ...MiddlewareHandler)  {
	s.HandleFunc(fmt.Sprintf("GET %s",route),hanler,middleware...)
}

func (s *RouterGroup) POST(route string, hanler RouteHandler, middleware ...MiddlewareHandler)  {
	s.HandleFunc(fmt.Sprintf("POST %s",route),hanler,middleware...)
}

func (s *RouterGroup) PATCH(route string, hanler RouteHandler, middleware ...MiddlewareHandler)  {
	s.HandleFunc(fmt.Sprintf("PATCH %s",route),hanler,middleware...)
}

func (s *RouterGroup) PUT(route string, hanler RouteHandler, middleware ...MiddlewareHandler)  {
	s.HandleFunc(fmt.Sprintf("PUT %s",route),hanler,middleware...)
}

func (s *RouterGroup) DELETE(route string, hanler RouteHandler, middleware ...MiddlewareHandler)  {
	s.HandleFunc(fmt.Sprintf("DELETE %s",route),hanler,middleware...)
}

func (s *RouterGroup) HEAD(route string, hanler RouteHandler, middleware ...MiddlewareHandler)  {
	s.HandleFunc(fmt.Sprintf("HEAD %s",route),hanler,middleware...)
}
