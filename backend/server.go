package backend

import (
	"net/http"
)

type Server struct {
	httpserver *http.Server
}

func (server *Server) Run(port string, handler http.Handler) error {
	server.httpserver = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	return server.httpserver.ListenAndServe()
}
