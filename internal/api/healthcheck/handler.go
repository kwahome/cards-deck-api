package healthcheck

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CheckStatusHandler represents a handler for healthcheck operations.
type CheckStatusHandler struct {
}

// NewCheckStatusHandler creates and returns a new handler.
//
//	returns: a CheckStatusHandler.
func NewCheckStatusHandler() CheckStatusHandler {
	return CheckStatusHandler{}
}

// Ping responds with a live-ness status.
func (handler *CheckStatusHandler) Ping(context *gin.Context) {
	response := StatusResponse{
		Status:    "alive",
		Timestamp: time.Now().Format(time.RFC850),
	}

	context.JSON(http.StatusOK, response)
}
