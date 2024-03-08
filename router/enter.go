package router

import (
	"server-fiber/router/app"
	"server-fiber/router/example"
	"server-fiber/router/frontend"
	"server-fiber/router/mobile"
	"server-fiber/router/order"
	"server-fiber/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Order    order.RouterGroup
	App      app.RouterGroup
	Frontend frontend.RouterGroup
	Mobile   mobile.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
