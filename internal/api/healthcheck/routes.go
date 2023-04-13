package healthcheck

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router gin.IRouter) {
	checkStatusHandler := NewCheckStatusHandler()
	router.GET("/healthcheck", checkStatusHandler.Ping)
}
