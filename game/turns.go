package game

import "jass/deck"

func ValidTurns(tableCards []deck.Card, handCards []deck.Card) []deck.Card {
	if len(tableCards) == 0 {
		return handCards[:]
	}

	leadingSuit := tableCards[0].Suit()
	cardsInCorrectSuit := filterBySuit(handCards, leadingSuit)
	if len(cardsInCorrectSuit) > 0 {
		return cardsInCorrectSuit
	}

	return handCards
}

func filterBySuit(cards []deck.Card, suit deck.Suit) []deck.Card{
	var filtered []deck.Card
	for _, c := range cards {
		if c.Suit() == suit {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func WinningPlay(plays deck.Plays) deck.Play {

	highestPlay := plays[0]
	leadingSuit := highestPlay.Card.Suit()
	for _, p := range plays[1:] {
		if p.Card.Suit() == leadingSuit && highestPlay.Card.CardValue() < p.Card.CardValue() {
			highestPlay = p
		}
	}

	return highestPlay
}