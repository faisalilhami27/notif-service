package routes

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"notification-service/routes/v1/messageLog"
	"notification-service/routes/v1/sms"
	"notification-service/routes/v1/whatsapp"
	"notification-service/src/controllers"
	"notification-service/utils"
	"os"
)

type Route struct {
	controller controllers.Controller
}

func NewRoute(controller controllers.Controller) *Route {
	return &Route{
		controller: controller,
	}
}

func (r *Route) Run() {
	router := gin.New()
	router.Use(cors.Default())
	router.Static("assets", "./public")
	group := router.Group("/api/v1")
	categoryRoute := messageLog.NewMessageLogRoute(r.controller.MessageLogController)
	categoryRoute.Start(group)

	whatsappRoute := whatsapp.NewWhatsappRoute(r.controller.WhatsappController)
	whatsappRoute.Start(group)

	smsRoute := sms.NewSMSRoute(r.controller.SMSController)
	smsRoute.Start(group)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	err := router.Run(port)
	if err != nil {
		utils.GetErrorLog(err)
		return
	}
}
