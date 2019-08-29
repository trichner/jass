package bots

import (
	"jass/deck"
	"math/rand"
)

type Player interface {
	PlayTurn(tableCards deck.Plays) deck.Card
}

type randomBot struct {
	handCards []deck.Card
}

func NewRandom(handCards []deck.Card) Player {
	return &randomBot{handCards:handCards}
}

func (r *randomBot) PlayTurn(tableCards deck.Plays) deck.Card {
	validPlays := deck.ValidPlays(tableCards.Cards(), r.handCards)

	i := rand.Intn(len(validPlays))
	play := validPlays[i]

	r.handCards, _ = deck.RemoveCard(r.handCards, play)
	return play
}


