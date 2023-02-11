package messageLog

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notification-service/src/usecases/v1/messageLog"
	"notification-service/utils"
)

type MessageLogController interface {
	GetAllCategory(ctx *gin.Context)
}

type MessageLogUseCase struct {
	useCase messageLog.MessageLogUseCase
}

func NewMessageLogController(useCase messageLog.MessageLogUseCase) *MessageLogUseCase {
	return &MessageLogUseCase{useCase}
}

func (c *MessageLogUseCase) GetAllCategory(ctx *gin.Context) {
	categories, err := c.useCase.GetAllCategory()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ReturnError(err))
		return
	}

	ctx.JSON(http.StatusOK, utils.ReturnSuccess(categories))
}
