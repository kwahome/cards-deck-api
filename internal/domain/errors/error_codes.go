package errors

const (
	// DeckNotFoundErrorCode error code for when the deck id provided cannot be found.
	DeckNotFoundErrorCode = "InvalidRequest.DeckNotFound"

	// DeckOutOfCardsErrorCode error code for when the deck is out of cards.
	DeckOutOfCardsErrorCode = "InvalidRequest.DeckOutOfCards"

	// InvalidCardErrorCode error code for when card is invalid.
	InvalidCardErrorCode = "InvalidRequest.InvalidCard"

	// InvalidCardsCountErrorCode error code for when count is invalid.
	InvalidCardsCountErrorCode = "InvalidRequest.InvalidCardsCount"

	// InvalidDeckIdErrorCode error code for when deck id is invalid.
	InvalidDeckIdErrorCode = "InvalidRequest.InvalidDeckId"

	// InternalServerErrorErrorCode error code for when an internal server error happens
	InternalServerErrorErrorCode = "ServerError.InternalServerError"
)
