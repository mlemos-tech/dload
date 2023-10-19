package user

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"mikaellemos.com.br/dload/src/model"
	"mikaellemos.com.br/dload/src/service"
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

	list := service.List(int(parseInt))
	utils.Render(ctx, 200, &list)
}

func Create(ctx *fasthttp.RequestCtx) {
	newUser, apiError := service.Create(ctx.Request.Body())

	if apiError != nil {
		utils.Render(ctx, 400, apiError)
		return
	}

	utils.Render(ctx, 201, &newUser)
}

func Update(ctx *fasthttp.RequestCtx) {
	param := ctx.QueryArgs().Peek("id")
	id, err := uuid.Parse(string(param))
	if err != nil {
		utils.Render(ctx, 400, &model.Message{"Invalid request"})
	}

	user := service.Update(id, ctx.Request.Body())
	utils.Render(ctx, 200, user)
}

func Delete(ctx *fasthttp.RequestCtx) {
	param := ctx.QueryArgs().Peek("id")
	id, err := uuid.Parse(string(param))

	if err != nil {
		logrus.Error(err)
		utils.Render(ctx, 400, &model.Message{"Invalid request"})
		return
	}

	service.Remove(id)
	utils.Render(ctx, 202, nil)
}

func Show(ctx *fasthttp.RequestCtx) {
	param := ctx.QueryArgs().Peek("id")
	id, err := uuid.Parse(string(param))

	if err != nil {
		logrus.Error(err)
		utils.Render(ctx, 400, &model.Message{"Invalid request"})
		return
	}

	user, _ := service.FindById(id)
	utils.Render(ctx, 200, user)
}
