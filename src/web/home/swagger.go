package home

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
	"mikaellemos.com.br/dload/src/config"
	"mikaellemos.com.br/dload/src/model"
	"mikaellemos.com.br/dload/src/service/user"
)

func Swagger(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("application/json")

	var data model.User
	client, _ := config.Connect("mongodb://pp:pp-payments@localhost:27017")
	db := client.Database("dload")

	list := data.List(context.Background(), db, "users", &data)
	payload, _ := json.Marshal(&list)

	fmt.Fprint(ctx, string(payload))

}

func PostSwagger(ctx *fasthttp.RequestCtx) {

	model, err := user.ValidateUser(ctx.Request.Body())
	client, _ := config.Connect("mongodb://pp:pp-payments@localhost:27017")

	if err != nil {

		body, _ := json.Marshal(&err)

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(400)

		ctx.SetBody(body)

		return
	}

	db := client.Database("dload")
	model.Create(context.Background(), db, "users", &model)

	ctx.SetContentType("application/json")
	fmt.Fprint(ctx, `{"message": "Ola mundo"}`)
}
