package sms

import (
	"github.com/gin-gonic/gin"
	"net/http"
	smsDT "notification-service/interfaces/dataTransfer/v1/sms"
	"notification-service/src/usecases/v1/sms"
	"notification-service/utils"
)

type SMSController interface {
	PublishMessage(ctx *gin.Context)
}

type SMSUseCase struct {
	useCase sms.SMSUseCase
}

func NewSMSController(useCase sms.SMSUseCase) *SMSUseCase {
	return &SMSUseCase{useCase}
}

func (c *SMSUseCase) PublishMessage(ctx *gin.Context) {
	var data smsDT.SMSBody
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
