package home

import (
	"github.com/valyala/fasthttp"
	"mikaellemos.com.br/dload/src/model"
	"mikaellemos.com.br/dload/src/web/utils"
)

func Index(ctx *fasthttp.RequestCtx) {
	utils.Render(ctx, 200, model.Message{"Hello World!"})
}
