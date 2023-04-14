package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kwahome/cards-deck-api/internal/api/v1/dtos"
	"github.com/kwahome/cards-deck-api/internal/api/v1/handlers"
	"github.com/kwahome/cards-deck-api/internal/domain/errors"
	"github.com/kwahome/cards-deck-api/internal/domain/model"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	httpHelpers "github.com/kwahome/cards-deck-api/pkg/http"
	"github.com/kwahome/cards-deck-api/tests/mocks"
	"github.com/kwahome/cards-deck-api/tests/unit"
	"github.com/stretchr/testify/assert"
)

func TestOpenDeck(t *testing.T) {

	t.Run("valid request should succeed", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockDeckService := mocks.NewMockDeckService(ctrl)

		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		deckId := "acee26ae-c304-4747-ab47-0109c6130a10"
		params := []gin.Param{
			{
				Key:   "id",
				Value: deckId,
			},
		}

		testhelpers.MockJsonGet(context, params, url.Values{})

		// stub
		mockDeckService.
			EXPECT().
			OpenDeck(deckId).
			DoAndReturn(func(string) (*model.Deck, error) {
				return &model.Deck{
					ID:        deckId,
					Shuffled:  false,
					Cards:     model.FullDeck[:],
					DrawCount: 52,
					Mutex:     sync.Mutex{},
				}, nil
			}).
			AnyTimes()

		handler := handlers.NewGetDeckHandler(mockDeckService)

		handler.OpenDeck(context)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response dtos.DeckOfCardsResponse
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.True(t, testhelpers.IsValidUUID(response.DeckResponse.DeckID))
		assert.False(t, response.DeckResponse.Shuffled)
		assert.Equal(t, 52, response.DeckResponse.Remaining)
		assert.Equal(t, 52, len(response.Cards))
	})

	t.Run("deck id supplied in request not found should return an error response", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		deckId := "acee26ae-c304-4747-ab47-0109c6130a10"
		params := []gin.Param{
			{
				Key:   "id",
				Value: deckId,
			},
		}

		testhelpers.MockJsonGet(context, params, url.Values{})

		handler := handlers.NewGetDeckHandler(service.CreateDeckService())

		handler.OpenDeck(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var response httpHelpers.ErrorResponse
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, errors.DeckNotFoundErrorCode, response.Code)
		assert.Equal(t, "the requested deck was not found", response.Message)
	})
}
