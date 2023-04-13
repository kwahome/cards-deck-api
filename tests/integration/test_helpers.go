package integration

import "github.com/gin-gonic/gin"

func TestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.UseRawPath = true
	return router
}
