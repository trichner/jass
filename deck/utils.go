package deck

func ValidPlays(tableCards []Card, handCards []Card) []Card {
	if len(tableCards) == 0 {
		return handCards[:]
	}

	leadingSuit := tableCards[0].Suit()
	cardsInCorrectSuit := CardsFilteredBySuit(handCards, leadingSuit)
	if len(cardsInCorrectSuit) > 0 {
		return cardsInCorrectSuit
	}

	return handCards
}

func CardsFilteredBySuit(cards []Card, suit Suit) []Card{
	var filtered []Card
	for _, c := range cards {
		if c.Suit() == suit {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func WinningCardIndex(tableCards []Card) int {

	highestInSuitIndex := 0
	highestInSuit := tableCards[highestInSuitIndex]
	leadingSuit := highestInSuit.Suit()
	for i, c := range tableCards[1:] {
		if c.Suit() == leadingSuit && highestInSuit.CardValue() < c.CardValue() {
			highestInSuitIndex = i
		}
	}

	return highestInSuitIndex
}

func ContainsCard(haystack []Card, needle Card) bool {
	for _, c := range haystack {
		if c == needle {
			return true
		}
	}

	return false
}

func RemoveCard(haystack []Card, needle Card) ([]Card, Card) {

	for i, c := range haystack {
		if c == needle {
			removed := haystack[i]
			haystack[i] = haystack[len(haystack)-1]
			return haystack[:len(haystack)-1], removed
		}
	}

	return haystack, None
}