package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kwahome/cards-deck-api/internal/api/v1/handlers"
	"github.com/kwahome/cards-deck-api/internal/domain/model"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	"github.com/kwahome/cards-deck-api/test/integration"
	"github.com/kwahome/cards-deck-api/test/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDrawCardsApi(t *testing.T) {

	const (
		deckId   = "196f3f6c-c738-4a74-a431-c0215c82b24e"
		endpoint = "/api/v1/decks/%s/draw?count=%s"
	)

	t.Run("should return the drawn cards when deck id and count are valid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		count := 2

		mockService := mocks.NewMockDeckService(ctrl)
		mockService.
			EXPECT().
			DrawCards(deckId, count).
			Return(model.FullDeck[0:count], nil)

		handler := handlers.NewDrawCardsHandler(mockService)

		router.GET("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(endpoint, deckId, strconv.Itoa(count)), nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)

		expectedResponse := `[{"code":"AC", "suite":"Club", "value":"Ace"}, {"code":"2C", "suite":"Club", "value":"2"}]`
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return bad request when deck id is invalid or could not be found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		count := 2

		handler := handlers.NewDrawCardsHandler(service.CreateDeckService())

		router.GET("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(endpoint, "invalid", strconv.Itoa(count)), nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		expectedResponse := `{"code":"InvalidRequest.DeckNotFound", "message":"the requested deck was not found"}`
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return bad request when count is invalid", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		handler := handlers.NewDrawCardsHandler(service.CreateDeckService())

		router.GET("/api/v1/decks/:id/draw", handler.DrawCards)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(endpoint, deckId, "invalid"), nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		expectedResponse := `{"code":"InvalidRequest.InvalidCardsCount", "message":"the cards count is invalid"}`
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}
