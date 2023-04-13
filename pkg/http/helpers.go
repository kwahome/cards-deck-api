package http

import (
	"github.com/gin-gonic/gin"
)

func RespondWithError(context *gin.Context, statusCode int, errorResponse ErrorResponse) {
	context.AbortWithStatusJSON(statusCode, errorResponse)
}
