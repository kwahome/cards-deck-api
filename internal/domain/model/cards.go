package model

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

type Cards []string

func (cards Cards) Shuffle() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	random.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
}

func (cards Cards) Validate() bool {
	for _, card := range cards {
		if !slices.Contains(FullDeck[:], card) {
			return false
		}
	}
	return true
}
