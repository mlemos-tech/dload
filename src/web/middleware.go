package web

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
	"mikaellemos.com.br/dload/src/service/userserv/validate"
	"mikaellemos.com.br/dload/src/web/utils"
)

func Middle(next phi.HandlerFunc) phi.HandlerFunc {
	return func(ctx *fasthttp.RequestCtx) {
		method := string(ctx.Method())

		if method == "POST" || method == "PUT" {
			_, err := validate.Validate(ctx.Request.Body())

			if err != nil {
				utils.Render(ctx, 403, &err)
				return
			}
		}

		next(ctx)
	}
}
