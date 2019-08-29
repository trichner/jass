package deck

var pointsTopdown = []int{
	0,
	0,
	8,
	0,
	10,
	2,
	3,
	4,
	11,
}

var pointsBottomup= []int{
	11,
	0,
	8,
	0,
	10,
	2,
	3,
	4,
	0,
}

var pointsTrumpTopDown = []int{
	0,
	0,
	0,
	0,
	10,
	2,
	3,
	4,
	11,
}

var pointsTrump = []int{
	0,
	0,
	0,
	14,
	10,
	20,
	3,
	4,
	11,
}

func PointsForTopDown(c Card) int {
	return pointsTopdown[c.CardValue()]
}

func PointsForBottomUp(c Card) int {
	return pointsBottomup[c.CardValue()]
}

func PointsForTrump(c Card, trumpSuit Suit) int {
	if c.Suit() == trumpSuit {
		return pointsTrump[c.CardValue()]
	}
	return pointsTrumpTopDown[c.CardValue()]
}