package server

import (
	"fmt"

	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	Host    string
	Port    int
	Handler phi.Handler
}

func (s Server) ListenAndServe() error {
	var g errgroup.Group

	g.Go(func() error {
		return fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", s.Host, s.Port), s.Handler.ServeFastHTTP)
	})

	return g.Wait()
}
