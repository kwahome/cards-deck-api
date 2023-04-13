package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwahome/cards-deck-api/internal/api/v1/helpers"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	"github.com/sirupsen/logrus"
)

// GetDeckHandler represents a handler for deck retrieval operations.
type GetDeckHandler struct {
	deckService service.DeckService
}

// NewGetDeckHandler creates and returns a new handler.
//
//	param: deckService - a DeckService
//
//	returns: a GetDeckHandler.
func NewGetDeckHandler(deckService service.DeckService) *GetDeckHandler {
	return &GetDeckHandler{deckService: deckService}
}

// OpenDeck retrieve a deck given its UUID.
//
// It accepts a deck id path param.
//
//	returns: a CreateDeckHandler.
//
// It returns the deck as a http response.
func (handler *GetDeckHandler) OpenDeck(context *gin.Context) {
	deckID := context.Param("id")

	logger := logrus.WithField("deckID", deckID)

	logger.Info("processing a request to retrieve a deck")

	deck, err := handler.deckService.OpenDeck(deckID)

	if err != nil {
		logger.Error(fmt.Sprintf("open deck request has failed due to: '%s'", err.Error()))
		helpers.GenerateAndRespondWithErrorResponse(context, err)
		return
	}

	context.JSON(http.StatusOK, helpers.GenerateDeckOfCardsResponse(deck))

	logger.Info("get deck request completed successfully")
}
