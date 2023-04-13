package dtos

type DeckOfCardsResponse struct {
	DeckResponse
	Cards []CardResponse `json:"cards"`
}
