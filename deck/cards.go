package deck

type CardValue int

//go:generate stringer -type=CardValue
const (
	Six CardValue = iota
	Seven
	Eight
	Nine
	Banner
	Under
	Ober
	King
	Ace
)

const suitOffset = 4

type Card uint8
const (
	None Card = 0xFF
	SchellenSix    Card = Card(uint8(Schellen)<<suitOffset | uint8(Six))
	SchellenSeven  Card = Card(uint8(Schellen)<<suitOffset | uint8(Seven))
	SchellenEight  Card = Card(uint8(Schellen)<<suitOffset | uint8(Eight))
	SchellenNine   Card = Card(uint8(Schellen)<<suitOffset | uint8(Nine))
	SchellenBanner Card = Card(uint8(Schellen)<<suitOffset | uint8(Banner))
	SchellenUnder  Card = Card(uint8(Schellen)<<suitOffset | uint8(Under))
	SchellenOber   Card = Card(uint8(Schellen)<<suitOffset | uint8(Ober))
	SchellenKing   Card = Card(uint8(Schellen)<<suitOffset | uint8(King))
	SchellenAce    Card = Card(uint8(Schellen)<<suitOffset | uint8(Ace))

	RosenSix    Card = Card(uint8(Rosen)<<suitOffset | uint8(Six))
	RosenSeven  Card = Card(uint8(Rosen)<<suitOffset | uint8(Seven))
	RosenEight  Card = Card(uint8(Rosen)<<suitOffset | uint8(Eight))
	RosenNine   Card = Card(uint8(Rosen)<<suitOffset | uint8(Nine))
	RosenBanner Card = Card(uint8(Rosen)<<suitOffset | uint8(Banner))
	RosenUnder  Card = Card(uint8(Rosen)<<suitOffset | uint8(Under))
	RosenOber   Card = Card(uint8(Rosen)<<suitOffset | uint8(Ober))
	RosenKing   Card = Card(uint8(Rosen)<<suitOffset | uint8(King))
	RosenAce    Card = Card(uint8(Rosen)<<suitOffset | uint8(Ace))

	SchiltenSix    Card = Card(uint8(Schilten)<<suitOffset | uint8(Six))
	SchiltenSeven  Card = Card(uint8(Schilten)<<suitOffset | uint8(Seven))
	SchiltenEight  Card = Card(uint8(Schilten)<<suitOffset | uint8(Eight))
	SchiltenNine   Card = Card(uint8(Schilten)<<suitOffset | uint8(Nine))
	SchiltenBanner Card = Card(uint8(Schilten)<<suitOffset | uint8(Banner))
	SchiltenUnder  Card = Card(uint8(Schilten)<<suitOffset | uint8(Under))
	SchiltenOber   Card = Card(uint8(Schilten)<<suitOffset | uint8(Ober))
	SchiltenKing   Card = Card(uint8(Schilten)<<suitOffset | uint8(King))
	SchiltenAce    Card = Card(uint8(Schilten)<<suitOffset | uint8(Ace))

	EichelSix    Card = Card(uint8(Eichel)<<suitOffset | uint8(Six))
	EichelSeven  Card = Card(uint8(Eichel)<<suitOffset | uint8(Seven))
	EichelEight  Card = Card(uint8(Eichel)<<suitOffset | uint8(Eight))
	EichelNine   Card = Card(uint8(Eichel)<<suitOffset | uint8(Nine))
	EichelBanner Card = Card(uint8(Eichel)<<suitOffset | uint8(Banner))
	EichelUnder  Card = Card(uint8(Eichel)<<suitOffset | uint8(Under))
	EichelOber   Card = Card(uint8(Eichel)<<suitOffset | uint8(Ober))
	EichelKing   Card = Card(uint8(Eichel)<<suitOffset | uint8(King))
	EichelAce    Card = Card(uint8(Eichel)<<suitOffset | uint8(Ace))
)

func (c Card) Suit() Suit {
	return Suit(c >> suitOffset)
}

func (c Card) CardValue() CardValue {
	return CardValue(0x0F & uint8(c))
}

func (c Card) String() string {
	if c == None {
		return "None"
	}
	return c.Suit().String() + c.CardValue().String()
}