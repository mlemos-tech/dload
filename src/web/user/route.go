package user

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"mikaellemos.com.br/dload/src/service/userserv"
	"strconv"
)

func List(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")
	page := ctx.QueryArgs().Peek("page")
	logrus.Info(page)
	_page := string(page)
	parseInt, err := strconv.ParseInt(_page, 10, 64)

	if err != nil {
		parseInt = 0
	}

	list := userserv.List(parseInt)
	payload, _ := json.Marshal(&list)
	fmt.Fprint(ctx, string(payload))
}

func Create(ctx *fasthttp.RequestCtx) {
	newUser, err := userserv.Create(ctx.Request.Body())
	ctx.SetContentType("application/json")

	if err != nil {
		body, _ := json.Marshal(&err)
		ctx.SetStatusCode(400)
		ctx.SetBody(body)

		return
	}

	body, _ := json.Marshal(newUser)
	ctx.SetStatusCode(201)
	ctx.SetBody(body)
}

func Update(ctx *fasthttp.RequestCtx) {

	ctx.SetContentType("application/json")
	id := ctx.QueryArgs().Peek("id")
	user, err := userserv.Update(string(id), ctx.Request.Body())

	if err != nil {
		body, _ := json.Marshal(&err)
		ctx.SetStatusCode(400)
		ctx.SetBody(body)

		return
	}

	body, _ := json.Marshal(user)
	ctx.SetStatusCode(201)
	ctx.SetBody(body)
}
