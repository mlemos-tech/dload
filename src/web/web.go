package web

import (
	"github.com/fate-lovely/phi"
	"mikaellemos.com.br/dload/src/web/home"
	"mikaellemos.com.br/dload/src/web/user"
)

func Handler() *phi.Mux {
	r := phi.NewRouter()

	r.Route("/users", func(r phi.Router) {
		r.Use(Middle)
		r.Get("/", user.List)
		r.Get("/show", user.Show)
		r.Post("/", user.Create)
		r.Put("/", user.Update)
		r.Delete("/", user.Delete)
	})

	r.Get("/", home.Index)

	return r
}
