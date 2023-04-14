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
	"github.com/kwahome/cards-deck-api/tests/integration"
	"github.com/kwahome/cards-deck-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCreateDeckApi(t *testing.T) {

	const (
		deckId   = "196f3f6c-c738-4a74-a431-c0215c82b24e"
		endpoint = "/api/v1/decks"
	)

	t.Run("should return a new full-deck when shuffle and cards params missing", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		mockService := mocks.NewMockDeckService(ctrl)

		mockService.
			EXPECT().
			CreateDeck(gomock.Any(), false).
			Return(
				&model.Deck{
					ID:        deckId,
					Shuffled:  false,
					Cards:     model.FullDeck[:],
					DrawCount: 52,
					Mutex:     sync.Mutex{},
				}, nil)

		handler := handlers.NewCreateDeckHandler(mockService)

		router.POST(endpoint, handler.CreateDeck)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, endpoint, nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Code)

		expectedResponse := fmt.Sprintf(`{"deck_id":"%s", "remaining":52, "shuffled":false}`, deckId)
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return a shuffled deck", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		mockService := mocks.NewMockDeckService(ctrl)
		mockService.
			EXPECT().
			CreateDeck(gomock.Any(), true).
			Return(
				&model.Deck{
					ID:        deckId,
					Shuffled:  true,
					Cards:     model.FullDeck[:],
					DrawCount: 52,
					Mutex:     sync.Mutex{},
				}, nil)

		handler := handlers.NewCreateDeckHandler(mockService)

		router.POST(endpoint, handler.CreateDeck)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?shuffle=%v", endpoint, true), nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Code)

		expectedResponse := fmt.Sprintf(`{"deck_id":"%s", "remaining":52, "shuffled":true}`, deckId)
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})

	t.Run("should return a shuffled deck with selected cards", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		router := integration.TestRouter()

		cards := model.Cards{"2H", "9D"}

		mockService := mocks.NewMockDeckService(ctrl)
		mockService.
			EXPECT().
			CreateDeck(gomock.Any(), true).
			Return(
				&model.Deck{
					ID:        deckId,
					Shuffled:  true,
					Cards:     cards,
					DrawCount: 52,
					Mutex:     sync.Mutex{},
				}, nil)

		handler := handlers.NewCreateDeckHandler(mockService)

		router.POST(endpoint, handler.CreateDeck)
		response := httptest.NewRecorder()
		request, _ := http.NewRequest(
			http.MethodPost,
			fmt.Sprintf("%s?shuffle=%v&cards=%s", endpoint, true, cards),
			nil)

		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusCreated, response.Code)

		expectedResponse := fmt.Sprintf(`{"deck_id":"%s", "remaining":2, "shuffled":true}`, deckId)
		assert.JSONEq(t, expectedResponse, response.Body.String())
	})
}
