package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kwahome/cards-deck-api/internal/api/v1/helpers"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	"github.com/sirupsen/logrus"
)

// CreateDeckHandler represents a handler for deck creation operations.
type CreateDeckHandler struct {
	deckService service.DeckService
}

// NewCreateDeckHandler creates and returns a new handler.
//
//	param: deckService - a DeckService
//
//	returns: a CreateDeckHandler.
func NewCreateDeckHandler(deckService service.DeckService) *CreateDeckHandler {
	return &CreateDeckHandler{deckService: deckService}
}

// CreateDeck create a new deck.
//
// It accepts shuffle and cards query params.
//
//	returns: a CreateDeckHandler.
//
// It returns the deck as a http response.
func (handler *CreateDeckHandler) CreateDeck(context *gin.Context) {
	shuffle := context.Query("shuffle") == "true"
	cardsParam := context.Query("cards")

	var cards []string
	if cardsParam != "" {
		cards = strings.Split(cardsParam, ",")
	}

	logger := logrus.WithField("cards", cards).WithField("shuffle", shuffle)

	logger.Info("processing a request to create a new deck")

	deck, err := handler.deckService.CreateDeck(cards, shuffle)
	if err != nil {
		logger.Error(fmt.Sprintf("the request to create a new deck has failed due to: '%s'", err.Error()))
		helpers.GenerateAndRespondWithErrorResponse(context, err)
		return
	}

	context.JSON(http.StatusCreated, helpers.GenerateDeckResponse(deck))

	logger.Info("create new deck request completed successfully")
}
