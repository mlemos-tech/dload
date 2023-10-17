package web

import (
	"github.com/fate-lovely/phi"
	"mikaellemos.com.br/dload/src/web/home"
)

func Handler() *phi.Mux {
	r := phi.NewRouter()

	r.Get("/", home.Index)
	r.Get("/swagger", home.Swagger)
	r.Post("/swagger", home.PostSwagger)

	return r
}
