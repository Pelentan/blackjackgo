package game

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit  string
	Face  string
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("|%s of %s|", c.Face, c.Suit)
}

type Deck []Card

var cardSuits = []string{"Hearts", "Diamonds", "Clubs", "Spades"}
var cardValues = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
var cardFaces = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}

func NewDeck() Deck {
	deck := make(Deck, 0)
	for _, suit := range cardSuits {
		for i, face := range cardFaces {
			deck = append(deck, Card{Suit: suit, Face: face, Value: cardValues[i]})
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
	} // I put this in as crash-protection.  But it's not tested for by the callers.
	card := (*d)[0]
	*d = (*d)[1:]
	return card
}
