package messageLog

import (
	"github.com/gin-gonic/gin"
	"notification-service/src/controllers/v1/messageLog"
)

type Route struct {
	router     *gin.Engine
	controller messageLog.MessageLogController
}

func NewMessageLogRoute(controller messageLog.MessageLogController) *Route {
	return &Route{
		router:     gin.Default(),
		controller: controller,
	}
}

func (r *Route) Start(route *gin.RouterGroup) *gin.Engine {
	messageLogRoute := route.Group("/message-log")
	messageLogRoute.GET("/", r.controller.GetAllCategory)
	return r.router
}
