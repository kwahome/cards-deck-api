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
	"github.com/kwahome/cards-deck-api/test/unit"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeck_Succeeds(t *testing.T) {

	recorder := httptest.NewRecorder()

	context := testhelpers.GetTestGinContext(recorder)

	testhelpers.MockJsonGet(context, []gin.Param{}, url.Values{})

	handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

	handler.CreateDeck(context)

	assert.Equal(t, http.StatusOK, recorder.Code)

	response := dtos.DeckResponse{}
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, testhelpers.IsValidUUID(response.DeckID))
	assert.False(t, response.Shuffled)
	assert.Equal(t, 52, response.Remaining)

}

func TestCreateDeck_WithShuffleQueryParam_Succeeds(t *testing.T) {

	recorder := httptest.NewRecorder()

	context := testhelpers.GetTestGinContext(recorder)

	queryParams := url.Values{}
	queryParams.Add("shuffle", "true")
	testhelpers.MockJsonGet(context, []gin.Param{}, queryParams)

	handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

	handler.CreateDeck(context)

	assert.Equal(t, http.StatusOK, recorder.Code)

	response := dtos.DeckResponse{}
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, testhelpers.IsValidUUID(response.DeckID))
	assert.True(t, response.Shuffled)
	assert.Equal(t, 52, response.Remaining)
}

func TestCreateDeck_WithCardsQueryParam_Succeeds(t *testing.T) {

	recorder := httptest.NewRecorder()

	context := testhelpers.GetTestGinContext(recorder)

	queryParams := url.Values{}
	queryParams.Add("cards", "AS,KD,AC,2C,KH")
	testhelpers.MockJsonGet(context, []gin.Param{}, queryParams)

	handler := handlers.NewCreateDeckHandler(service.CreateDeckService())

	handler.CreateDeck(context)

	assert.Equal(t, http.StatusOK, recorder.Code)

	response := dtos.DeckResponse{}
	err := json.Unmarshal([]byte(recorder.Body.String()), &response)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, testhelpers.IsValidUUID(response.DeckID))
	assert.False(t, response.Shuffled)
	assert.Equal(t, 5, response.Remaining)
}
