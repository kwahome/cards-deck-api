package helpers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwahome/cards-deck-api/internal/api/v1/dtos"
	serviceErrors "github.com/kwahome/cards-deck-api/internal/domain/errors"
	"github.com/kwahome/cards-deck-api/internal/domain/model"
	httpHelpers "github.com/kwahome/cards-deck-api/pkg/http"
)

func GenerateCardsResponse(cards model.Cards) []dtos.CardResponse {

	var cardsResponse []dtos.CardResponse
	for _, card := range cards {
		cardsResponse = append(cardsResponse, dtos.CardResponse{
			Value: model.CardSymbolNameMap[card[:len(card)-1]],
			Suite: model.CardSuitMap[card[len(card)-1:]],
			Code:  card,
		})
	}
	return cardsResponse
}

func GenerateDeckResponse(deck *model.Deck) *dtos.DeckResponse {
	return &dtos.DeckResponse{
		DeckID:    deck.ID,
		Shuffled:  deck.Shuffled,
		Remaining: len(deck.Cards),
	}
}

func GenerateDeckOfCardsResponse(deckModel *model.Deck) *dtos.DeckOfCardsResponse {
	deck := GenerateDeckResponse(deckModel)

	cards := GenerateCardsResponse(deckModel.Cards)

	response := &dtos.DeckOfCardsResponse{
		DeckResponse: *deck,
		Cards:        cards,
	}

	return response
}

func GenerateAndRespondWithErrorResponse(context *gin.Context, err error) {
	var serviceError *serviceErrors.ServiceError
	errors.As(err, &serviceError)

	var errorResponse httpHelpers.ErrorResponse
	var httpStatusCode int
	if serviceError != nil {
		errorResponse = httpHelpers.ErrorResponse{
			Code:    serviceError.Code,
			Message: serviceError.Message,
		}

		httpStatusCode = http.StatusBadRequest
	} else {
		errorResponse = httpHelpers.ErrorResponse{
			Code:    serviceErrors.InternalServerErrorErrorCode,
			Message: "an internal server error has occurred",
		}

		httpStatusCode = http.StatusInternalServerError
	}

	httpHelpers.RespondWithError(context, httpStatusCode, errorResponse)
}
