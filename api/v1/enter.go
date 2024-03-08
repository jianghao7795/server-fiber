package v1

import (
	"server-fiber/api/v1/app"
	"server-fiber/api/v1/example"
	"server-fiber/api/v1/frontend"
	"server-fiber/api/v1/mobile"
	"server-fiber/api/v1/order"
	"server-fiber/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	OrderApiGroup    order.ApiGroup
	AppApiGroup      app.ApiGroup
	FrontendApiGroup frontend.ApiGroup
	MobileApiGroup   mobile.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
