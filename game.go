package main

import (
	"fmt"
	"jass/deck"
	"jass/game"
	"math/rand"
	"time"
)

func main() {
	g := game.NewGame()
	for i := 0;i<9;i++ {
		g.NextRound()
	}

	for _, t := range g.Tricks() {
		fmt.Println(t.String())
	}
}

func newShuffledDeck() []deck.Card {
	cards := deck.NewDeck()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rnd.Shuffle(len(cards), func(i, j int) {
		t := cards[i]
		cards[i] = cards[j]
		cards[j] = t
	})
	return cards
}

