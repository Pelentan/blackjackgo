package game

import (
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Value int
}

type Deck []Card

var cardSuits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var cardValues = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}

func NewDeck() Deck {
	deck := make(Deck, 0)
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			deck = append(deck, Card{Suit: suit, Value: value})
		}
	}
	return deck
}

// Creates a new random source with a seed based on the current time and shuffles the deck.
func (d *Deck) ShuffleDeck() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	rng.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d *Deck) DealCard() Card {
	if len(*d) == 0 {
		return Card{}
	}
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
