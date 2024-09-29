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
