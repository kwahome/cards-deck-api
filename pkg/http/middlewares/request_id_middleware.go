package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kwahome/cards-deck-api/pkg/http"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestId := context.GetHeader(http.RequestId)

		if len(requestId) == 0 {
			context.Writer.Header().Set(http.RequestId, uuid.NewString())
		}
		context.Next()
	}
}
