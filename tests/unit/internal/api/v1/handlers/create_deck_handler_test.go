package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kwahome/cards-deck-api/internal/api/v1/dtos"
	"github.com/kwahome/cards-deck-api/internal/api/v1/handlers"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	"github.com/kwahome/cards-deck-api/tests/unit"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {

	t.Run("valid request should succeed", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		testhelpers.MockJsonGet(context, []gin.Param{}, url.Values{})

		handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

		handler.CreateDeck(context)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		response := dtos.DeckResponse{}
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.True(t, testhelpers.IsValidUUID(response.DeckID))
		assert.False(t, response.Shuffled)
		assert.Equal(t, 52, response.Remaining)
	})

	t.Run("valid request with shuffle query param should succeed", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		queryParams := url.Values{}
		queryParams.Add("shuffle", "true")
		testhelpers.MockJsonGet(context, []gin.Param{}, queryParams)

		handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

		handler.CreateDeck(context)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		response := dtos.DeckResponse{}
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.True(t, testhelpers.IsValidUUID(response.DeckID))
		assert.True(t, response.Shuffled)
		assert.Equal(t, 52, response.Remaining)
	})

	t.Run("valid request with cards query param should succeed", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		queryParams := url.Values{}
		queryParams.Add("cards", "AS,KD,AC,2C,KH")
		testhelpers.MockJsonGet(context, []gin.Param{}, queryParams)

		handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

		handler.CreateDeck(context)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		response := dtos.DeckResponse{}
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.True(t, testhelpers.IsValidUUID(response.DeckID))
		assert.False(t, response.Shuffled)
		assert.Equal(t, 5, response.Remaining)
	})

	t.Run("valid request with shuffle and cards query params should succeed", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		queryParams := url.Values{}
		queryParams.Add("shuffle", "true")
		queryParams.Add("cards", "AS,KD,AC,2C,KH")
		testhelpers.MockJsonGet(context, []gin.Param{}, queryParams)

		handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

		handler.CreateDeck(context)

		assert.Equal(t, http.StatusCreated, recorder.Code)

		response := dtos.DeckResponse{}
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.True(t, testhelpers.IsValidUUID(response.DeckID))
		assert.True(t, response.Shuffled)
		assert.Equal(t, 5, response.Remaining)
	})
}
