package controller

import (
	"github.com/mizuki1412/go-core-kit/service/restkit/context"
	"github.com/mizuki1412/go-core-kit/service/restkit/router"
	"wechat-work-pusher/middleware"
	"wechat-work-pusher/service"
)

func Init(router *router.Router) {
	r := router.Use(middleware.AuthToken())
	{
		r.Post("/msg", msg)
		r.Post("/card", card)
	}
}

type msgParams struct {
	To      string
	Content string `validate:"required"`
}

func msg(ctx *context.Context) {
	params := msgParams{}
	ctx.BindForm(&params)
	service.SendMsg(params.To, params.Content)
	ctx.JsonSuccess("ok")
}

type cardParams struct {
	To          string
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Url         string `validate:"required"`
}

func card(ctx *context.Context) {
	params := cardParams{}
	ctx.BindForm(&params)
	service.SendCardMsg(params.To, params.Title, params.Description, params.Url)
	ctx.JsonSuccess("ok")
}
