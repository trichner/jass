package deck

type Suit uint8

//go:generate stringer -type=Suit
const (
	Schellen Suit = iota
	Rosen
	Schilten
	Eichel
)
