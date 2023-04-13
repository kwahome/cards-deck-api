package service

import (
	"sync"

	"github.com/google/uuid"
	"github.com/kwahome/cards-deck-api/internal/domain/errors"
	"github.com/kwahome/cards-deck-api/internal/domain/model"
)

// DeckService is a service for interacting with a deck.
type DeckService interface {
	CreateDeck(cards model.Cards, shuffle bool) (*model.Deck, error)
	DrawCards(deckID string, count int) (model.Cards, error)
	OpenDeck(deckID string) (*model.Deck, error)
}

type deckService struct {
}

// CreateDeckService returns a new instance of DeckService.
func CreateDeckService() DeckService {
	return &deckService{}
}

// CreateDeck creates a new deck.
//
//	param:	cards		- the cards in the deck
//	param:	shuffle		- whether the deck should be shuffled
//
//	returns: the newly created deck and any error that occurred during the operation.
func (service *deckService) CreateDeck(cards model.Cards, shuffle bool) (*model.Deck, error) {
	if !cards.Validate() {
		return nil, errors.NewServiceError(errors.InvalidCardErrorCode, "invalid card provided")
	}

	deckID := uuid.NewString()

	deck := model.Deck{
		ID:       deckID,
		Shuffled: shuffle,
		Cards:    createCards(shuffle, cards),
	}

	model.Decks[deckID] = &deck

	return &deck, nil
}

// DrawCards draws cards from a deck.
//
//	param:	deckID	- the id of the deck to draw from
//	param:	count	- the number of cards to draw
//
//	returns: the drawn cards and any error that occurred during the operation.
func (service *deckService) DrawCards(deckID string, count int) (model.Cards, error) {
	deck, exist := model.Decks[deckID]
	if !exist {
		return nil, errors.NewServiceError(errors.DeckNotFoundErrorCode, "the requested deck was not found")
	}

	if count+deck.DrawCount > len(deck.Cards) {
		return nil, errors.NewServiceError(errors.DeckOutOfCardsErrorCode, "the deck is out of cards")
	}

	// synchronize updates to draw count
	deck.Lock()
	deck.DrawCount += count
	deck.Unlock()

	return deck.Cards[deck.DrawCount-count : deck.DrawCount], nil
}

// OpenDeck retrieves a deck by its id.
//
//	param:	deckID - the id of the deck
//
//	returns: the deck and any error that occurred during the operation.
func (service *deckService) OpenDeck(deckID string) (*model.Deck, error) {
	deck, exist := model.Decks[deckID]
	if !exist {
		return nil, errors.NewServiceError(errors.DeckNotFoundErrorCode, "the requested deck was not found")
	}
	return &model.Deck{
		ID:        deck.ID,
		Shuffled:  deck.Shuffled,
		Cards:     deck.Cards[deck.DrawCount:],
		DrawCount: deck.DrawCount,
		Mutex:     sync.Mutex{},
	}, nil
}

func createCards(shuffle bool, cards model.Cards) model.Cards {
	if len(cards) == 0 {
		cards = model.FullDeck[:]
	}

	if shuffle {
		cards.Shuffle()
	}

	return cards
}
