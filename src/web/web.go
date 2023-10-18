package web

import (
	"github.com/fate-lovely/phi"
	"mikaellemos.com.br/dload/src/web/home"
	"mikaellemos.com.br/dload/src/web/user"
)

func Handler() *phi.Mux {
	r := phi.NewRouter()

	r.Get("/", home.Index)
	r.Get("/swagger", user.List)
	r.Post("/swagger", user.Create)
	r.Put("/swagger/{$}", user.Update)

	return r
}
