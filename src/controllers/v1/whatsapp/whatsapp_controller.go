package whatsapp

import (
	"github.com/gin-gonic/gin"
	"net/http"
	whatsapp2 "notification-service/interfaces/dataTransfer/v1/whatsapp"
	"notification-service/src/usecases/v1/whatsapp"
	"notification-service/utils"
)

type WhatsappController interface {
	PublishMessage(ctx *gin.Context)
}

type WhatsappUseCase struct {
	useCase whatsapp.WhatsappUseCase
}

func NewWhatsappController(useCase whatsapp.WhatsappUseCase) *WhatsappUseCase {
	return &WhatsappUseCase{useCase}
}

func (c *WhatsappUseCase) PublishMessage(ctx *gin.Context) {
	var data whatsapp2.WhatsappBody
	err := ctx.ShouldBindJSON(&data)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	err = c.useCase.PublishMessage(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(nil))
}
