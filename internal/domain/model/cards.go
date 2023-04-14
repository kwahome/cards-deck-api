package model

import (
	"math/rand"
	"time"

	"golang.org/x/exp/slices"
)

var CardSymbolNameMap = map[string]string{
	"A":  "ACE",
	"1":  "1",
	"2":  "2",
	"3":  "3",
	"4":  "4",
	"5":  "5",
	"6":  "6",
	"7":  "7",
	"8":  "8",
	"9":  "9",
	"10": "10",
	"J":  "JACK",
	"Q":  "QUEEN",
	"K":  "KING",
}

var CardSuitMap = map[string]string{
	"C": "CLUBS",
	"D": "DIAMONDS",
	"H": "HEARTS",
	"S": "SPADES",
}

var CardSuitSymbolMap = map[string]string{
	"C": "♣",
	"D": "♦",
	"H": "♥",
	"S": "♠",
}

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
