package whatsapp

import (
	"github.com/gin-gonic/gin"
	"notification-service/src/controllers/v1/whatsapp"
)

type Route struct {
	router     *gin.Engine
	controller whatsapp.WhatsappController
}

func NewWhatsappRoute(controller whatsapp.WhatsappController) *Route {
	return &Route{
		router:     gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	messageLogRoute := route.Group("/whatsapp/send")
	messageLogRoute.POST("/", r.controller.PublishMessage)
	return r.router
}
