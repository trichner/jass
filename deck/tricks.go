package deck

import (
	"fmt"
	"strings"
)

type PlayerIdentifier interface {
	ID() int
}

type Play struct {
	Card   Card
	Player PlayerIdentifier
}

func (p *Play) String() string {
	return fmt.Sprintf("(%s played by %d)", p.Card.String(), p.Player.ID())
}

type Plays []Play

func (p Plays) Cards() []Card {
	var cards []Card
	for _, play := range p {
		cards = append(cards, play.Card)
	}
	return cards
}

type Trick struct {
	plays Plays
	owner PlayerIdentifier
}

func NewTrick(owner PlayerIdentifier, plays Plays) *Trick {
	return &Trick{
		plays: plays,
		owner: owner,
	}
}

func (t *Trick) Cards() []Card {
	var cards []Card
	for _, p := range t.plays {
		cards = append(cards, p.Card)
	}
	return t.plays.Cards()
}

func (t *Trick) Plays() Plays {
	return t.plays
}

func (t *Trick) Owner() PlayerIdentifier {
	return t.owner
}

func (t *Trick) String() string {
	s := strings.Builder{}
	for _, c := range t.plays {
		s.WriteString(c.String())
		s.WriteString(" ")
	}

	s.WriteString(fmt.Sprintf("(%d)", t.owner.ID()))
	return s.String()
}
