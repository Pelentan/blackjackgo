package game

type Hand []Card

func (h Hand) CountCards() int {
	score := 0
	aceCount := 0

	for _, card := range h {
		if card.Value == 11 {
			aceCount++
		}
		score += card.Value
	}

	// Adjust for aces if the score is over 21
	for score > 21 && aceCount > 0 {
		score -= 10
		aceCount--
	}

	return score
}
