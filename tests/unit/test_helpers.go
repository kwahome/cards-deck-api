package testhelpers

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func IsValidUUID(input string) bool {
	_, err := uuid.Parse(input)
	return err == nil
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MockJsonGet(context *gin.Context, params gin.Params, u url.Values) {
	context.Request.Method = "GET"
	context.Request.Header.Set("Content-Type", "application/json")

	// set path params
	context.Params = params

	// set query params
	context.Request.URL.RawQuery = u.Encode()
}
