package healthcheck

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kwahome/cards-deck-api/internal/api/healthcheck"
	"github.com/kwahome/cards-deck-api/test/integration"
	"github.com/stretchr/testify/assert"
)

func TestGetHealthCheckApi(t *testing.T) {
	t.Run("should return live-ness status", func(t *testing.T) {
		endpoint := "/health"

		router := integration.TestRouter()

		healthCheckHandler := healthcheck.NewCheckStatusHandler()
		router.GET(endpoint, healthCheckHandler.Ping)

		// call the code we are testing
		request := httptest.NewRequest(http.MethodGet, endpoint, nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, request)

		response := healthcheck.StatusResponse{}

		assert.Equal(t, http.StatusOK, recorder.Code)

		err := json.NewDecoder(recorder.Body).Decode(&response)
		if err != nil {
			t.Error(err)
		}

		// assert that the expectations were met
		assert.Equal(t, "alive", response.Status)
	})
}
