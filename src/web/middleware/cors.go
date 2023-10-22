package middleware

import (
	"github.com/fate-lovely/phi"
	"github.com/valyala/fasthttp"
)

func Cors(next phi.HandlerFunc) phi.HandlerFunc {
	return func(ctx *fasthttp.RequestCtx) {
		method := string(ctx.Method())

		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "false")
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,OPTIONS")

		if method == "OPTIONS" {
			return
		}

		next(ctx)
	}
}
