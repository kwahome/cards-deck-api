package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kwahome/cards-deck-api/internal/api/v1/handlers"
	"github.com/kwahome/cards-deck-api/internal/domain/model"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	"github.com/kwahome/cards-deck-api/tests/integration"
	"github.com/kwahome/cards-deck-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestOpenDeckApi(t *testing.T) {

	const (
		deckId   = "196f3f6c-c738-4a74-a431-c0215c82b24e"
		endpoint = "/api/v1/decks/%s"
	)

	t.Run("should return the requested deck", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		mockService := mocks.NewMockDeckService(ctrl)
		mockService.
			EXPECT().
			OpenDeck(deckId).
			Return(
				&model.Deck{
					ID:        deckId,
					Shuffled:  false,
					Cards:     model.FullDeck[:2],
					DrawCount: 2,
					Mutex:     sync.Mutex{},
				}, nil)

		handler := handlers.NewGetDeckHandler(mockService)

		router.GET("/api/v1/decks/:id", handler.OpenDeck)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(endpoint, deckId), nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)

		expectedResponse := fmt.Sprintf(
			`{"deck_id":"%s", "remaining":2, "shuffled":false,
			"cards":[{"code":"AC", "suite":"Club", "value":"Ace"}, 
			{"code":"2C", "suite":"Club", "value":"2"}]}`,
			deckId)

		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return bad request when deck id is invalid or could not be found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		handler := handlers.NewGetDeckHandler(service.CreateDeckService())

		router.GET("/api/v1/decks/:id", handler.OpenDeck)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(endpoint, "invalid"), nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code)

		expectedResponse := `{"code":"InvalidRequest.DeckNotFound", "message":"the requested deck was not found"}`
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}
