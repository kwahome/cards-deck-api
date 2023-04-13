package healthcheck

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kwahome/cards-deck-api/internal/api/healthcheck"
	"github.com/kwahome/cards-deck-api/tests/unit"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckPing(t *testing.T) {

	recorder := httptest.NewRecorder()

	context := testhelpers.GetTestGinContext(recorder)

	testhelpers.MockJsonGet(context, []gin.Param{}, url.Values{})

	handler := healthcheck.CheckStatusHandler{}

	handler.Ping(context)

	assert.Equal(t, http.StatusOK, recorder.Code)

	statusResponse := healthcheck.StatusResponse{}
	err := json.Unmarshal([]byte(recorder.Body.String()), &statusResponse)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "alive", statusResponse.Status)

}
