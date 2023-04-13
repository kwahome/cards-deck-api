package middlewares

import (
	http2 "github.com/kwahome/cards-deck-api/pkg/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(authKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader(http2.AuthToken)

		if len(token) == 0 || token != authKey {
			errorResponse := http2.ErrorResponse{
				Code:    "InvalidRequest.Unauthorized",
				Message: "the authorization key is missing of invalid",
			}

			http2.RespondWithError(context, http.StatusUnauthorized, errorResponse)

			return
		}

		context.Next()
	}
}
