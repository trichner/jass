package game

import (
	"fmt"
	"jass/bots"
	"jass/deck"
)

type Game struct {
	players        []*playerState
	tricks         []*deck.Trick
	startingPlayer deck.PlayerIdentifier
}

func NewGame() *Game {
	players := make([]*playerState, 4)
	cards := deck.NewShuffledDeck()
	for i := 0; i < 4; i++ {
		start := i * 9
		end := (i + 1) * 9
		if end > len(cards) {
			end = len(cards)
		}
		players[i] = NewPlayer(i, cards[start:end])
	}
	g := &Game{
		players:        players,
		tricks:         []*deck.Trick{},
		startingPlayer: players[0],
	}
	return g
}

func (g *Game) NextRound() {
	fmt.Printf(">> round\n")

	tablePlays := make(deck.Plays, 0, 4)
	for i := 0; i < 4; i++ {

		player := g.currentPlayer(i)
		copiedTableCards := append(deck.Plays{}, tablePlays...)
		card := player.player.PlayTurn(copiedTableCards)

		play := g.playCard(player, tablePlays, card)
		tablePlays = append(tablePlays, play)
	}

	winningPlay := WinningPlay(tablePlays)
	trick := deck.NewTrick(winningPlay.Player, tablePlays)

	g.tricks = append(g.tricks, trick)
	g.startingPlayer = winningPlay.Player
}

type PlayerScore struct {
	Player deck.PlayerIdentifier
	Score int
}

func (g *Game) Score() []PlayerScore {

	for _, t := range g.tricks {
		t.Cards()
	}

}

func (g *Game) currentPlayer(i int) *playerState {
	playerIndex := (i + g.startingPlayer.ID()) % 4
	return g.players[playerIndex]
}

func (g *Game) playCard(player *playerState, tablePlays deck.Plays, card deck.Card) deck.Play {
	validTurns := ValidTurns(tablePlays.Cards(), player.cards)

	if !deck.ContainsCard(validTurns, card) {
		panic(fmt.Errorf("illegal play from player %d", player.ID()))
	}

	err := player.playCard(card)
	if err != nil {
		panic(fmt.Errorf("illegal play from player %d: %s", player.ID(), err))
	}
	return deck.Play{
		Card:   card,
		Player: player,
	}
}

func (g *Game) Tricks() []*deck.Trick {
	return g.tricks
}

type playerState struct {
	id     int
	cards  []deck.Card
	player bots.Player
}

func (p *playerState) ID() int {
	return p.id
}

func (p *playerState) validPlays(tableCards []deck.Card) []deck.Card {
	return ValidTurns(tableCards, p.cards)
}

func (p *playerState) playCard(card deck.Card) error {
	cards := p.cards
	cards, playedCard := deck.RemoveCard(cards, card)
	if playedCard != card {
		return fmt.Errorf("cannot play card %s, not found in hand cards", card)
	}
	p.cards = cards
	return nil
}

func NewPlayer(id int, cards []deck.Card) *playerState {

	p := &playerState{
		id:     id,
		cards:  cards,
		player: bots.NewRandom(append([]deck.Card{}, cards...)),
	}

	return p
}
