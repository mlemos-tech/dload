package user

import (
	"github.com/valyala/fasthttp"
	"mikaellemos.com.br/dload/src/service/userserv"
	"mikaellemos.com.br/dload/src/web/utils"
	"strconv"
)

func List(ctx *fasthttp.RequestCtx) {
	page := ctx.QueryArgs().Peek("page")
	_page := string(page)
	parseInt, err := strconv.ParseInt(_page, 10, 64)

	if err != nil {
		parseInt = 0
	}

	list := userserv.List(int(parseInt))
	utils.Render(ctx, 200, &list)
}

func Create(ctx *fasthttp.RequestCtx) {
	newUser, err := userserv.Create(ctx.Request.Body())

	if err != nil {
		utils.Render(ctx, 400, &err)
		return
	}

	utils.Render(ctx, 201, &newUser)
}

func Update(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")
	iduint, _ := strconv.ParseUint(string(id), 10, 64)
	user := userserv.Update(iduint, ctx.Request.Body())
	utils.Render(ctx, 200, user)
}

func Delete(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")
	iduint, _ := strconv.ParseUint(string(id), 10, 64)
	userserv.Remove(int(iduint))
	utils.Render(ctx, 202, nil)
}
