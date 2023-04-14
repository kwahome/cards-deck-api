package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
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

func TestDrawCardsFromDeck(t *testing.T) {

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

		cardCount := 5
		queryParams := url.Values{}
		queryParams.Add("count", strconv.Itoa(cardCount))
		testhelpers.MockJsonGet(context, params, queryParams)

		// stub
		mockDeckService.
			EXPECT().
			DrawCards(deckId, cardCount).
			DoAndReturn(func(string, int) (model.Cards, error) {
				return model.FullDeck[0:cardCount], nil
			}).
			AnyTimes()

		handler := handlers.NewDrawCardsHandler(mockDeckService)

		handler.DrawCards(context)

		assert.Equal(t, http.StatusOK, recorder.Code)

		var response []dtos.CardResponse
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, cardCount, len(response))
	})

	t.Run("missing or invalid deck id in request should return an error response", func(t *testing.T) {
		recorder := httptest.NewRecorder()

		context := testhelpers.GetTestGinContext(recorder)

		testhelpers.MockJsonGet(context, []gin.Param{}, url.Values{})

		handler := handlers.NewDrawCardsHandler(service.CreateDeckService())

		handler.DrawCards(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var response httpHelpers.ErrorResponse
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, errors.InvalidDeckIdErrorCode, response.Code)
		assert.Equal(t, "the deck id is invalid", response.Message)
	})

	t.Run("missing or invalid count in request should return an error response", func(t *testing.T) {
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

		handler := handlers.NewDrawCardsHandler(service.CreateDeckService())

		handler.DrawCards(context)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)

		var response httpHelpers.ErrorResponse
		err := json.Unmarshal([]byte(recorder.Body.String()), &response)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, errors.InvalidCardsCountErrorCode, response.Code)
		assert.Equal(t, "the cards count is invalid", response.Message)
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

		cardCount := 5
		queryParams := url.Values{}
		queryParams.Add("count", strconv.Itoa(cardCount))
		testhelpers.MockJsonGet(context, params, queryParams)

		handler := handlers.NewDrawCardsHandler(service.CreateDeckService())

		handler.DrawCards(context)

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
