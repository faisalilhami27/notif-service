package sms

import (
	"github.com/gin-gonic/gin"
	"notification-service/src/controllers/v1/sms"
)

type Route struct {
	router     *gin.Engine
	controller sms.SMSController
}

func NewSMSRoute(controller sms.SMSController) *Route {
	return &Route{
		router:     gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	messageLogRoute := route.Group("/sms/send")
	messageLogRoute.POST("/", r.controller.PublishMessage)
	return r.router
}
