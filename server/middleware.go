package server

type MiddlewareHandler func(next RouteHandler) RouteHandler

type Middleware interface{
	Use(handlers ...MiddlewareHandler) Middleware
}

func (s *Server) Use(handlers ...MiddlewareHandler) Middleware  {	
	if len(handlers) == 0 {
		panic("Middleware need to be specified")
	}

	s.globalMiddlewares = append(s.globalMiddlewares, handlers...)
	return s
}