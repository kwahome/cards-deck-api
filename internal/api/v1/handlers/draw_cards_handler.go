package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/kwahome/cards-deck-api/internal/api/v1/helpers"
	"github.com/kwahome/cards-deck-api/internal/domain/errors"
	"github.com/kwahome/cards-deck-api/internal/domain/service"
	"github.com/sirupsen/logrus"
)

// IDrawCardsHandler interface.
type IDrawCardsHandler interface {
	DrawCards(context *gin.Context)
}

// DrawCardsHandler represents a handler for card draw operations.
type DrawCardsHandler struct {
	deckService service.DeckService
}

// NewDrawCardsHandler creates and returns a new handler.
//
//	param: deckService - a DeckService
//
//	returns: a DrawCardsHandler.
func NewDrawCardsHandler(deckService service.DeckService) *DrawCardsHandler {
	return &DrawCardsHandler{deckService: deckService}
}

// DrawCards draws cards from the requested deck.
//
// It accepts a deck id path param and card count query param.
//
// It returns the drawn cards.
func (handler *DrawCardsHandler) DrawCards(context *gin.Context) {
	deckID := context.Param("id")
	if err := validation.Validate(deckID, validation.Required); err != nil {
		serviceError := errors.NewServiceError(errors.InvalidDeckIdErrorCode, "the deck id is invalid")
		helpers.GenerateAndRespondWithErrorResponse(context, serviceError)
		return
	}

	countParam := context.Query("count")
	var countParamRegex = regexp.MustCompile("^[0-9]{1,2}$")
	if err := validation.Validate(countParam, validation.Required, validation.Match(countParamRegex)); err != nil {
		serviceError := errors.NewServiceError(errors.InvalidCardsCountErrorCode, "the cards count is invalid")
		helpers.GenerateAndRespondWithErrorResponse(context, serviceError)
		return
	}

	count, _ := strconv.Atoi(countParam)

	logger := logrus.WithField("deckID", deckID).WithField("count", count)

	logger.Info("processing a request to draw cards")

	cards, err := handler.deckService.DrawCards(deckID, count)
	if err != nil {
		logger.Error(fmt.Sprintf("draw cards request has failed due to: '%s'", err.Error()))
		helpers.GenerateAndRespondWithErrorResponse(context, err)
		return
	}

	context.JSON(http.StatusOK, helpers.GenerateCardsResponse(cards))

	logger.Info("request to draw cards completed successfully")
}
