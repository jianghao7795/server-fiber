package service

import (
	"server-fiber/service/app"
	"server-fiber/service/example"
	"server-fiber/service/frontend"
	"server-fiber/service/mobile"
	"server-fiber/service/order"
	"server-fiber/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	OrderServiceGroup    order.ServiceGroup
	AppServiceGroup      app.ServiceGroup
	FrontendServiceGroup frontend.ServiceGroup
	MobileServiceGroup   mobile.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
