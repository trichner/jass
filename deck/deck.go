package deck

import (
	"math/rand"
	"time"
)

var deck = []Card{
	SchellenSix,
	SchellenSeven,
	SchellenEight,
	SchellenNine,
	SchellenBanner,
	SchellenUnder,
	SchellenOber,
	SchellenKing,
	SchellenAce,
	RosenSix,
	RosenSeven,
	RosenEight,
	RosenNine,
	RosenBanner,
	RosenUnder,
	RosenOber,
	RosenKing,
	RosenAce,
	SchiltenSix,
	SchiltenSeven,
	SchiltenEight,
	SchiltenNine,
	SchiltenBanner,
	SchiltenUnder,
	SchiltenOber,
	SchiltenKing,
	SchiltenAce,
	EichelSix,
	EichelSeven,
	EichelEight,
	EichelNine,
	EichelBanner,
	EichelUnder,
	EichelOber,
	EichelKing,
	EichelAce,
}

func NewDeck() []Card {
	return deck[:]
}

func NewShuffledDeck() []Card {
	cards := NewDeck()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd.Shuffle(len(cards), func(i, j int) {
		t := cards[i]
		cards[i] = cards[j]
		cards[j] = t
	})
	return cards
}