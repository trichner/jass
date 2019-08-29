package deck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCard_String(t *testing.T) {
	assert.Equal(t, "SchellenOber", SchellenOber.String(), "Card names should be correct")
	assert.Equal(t, "None", None.String(), "Card names should be correct")
}

func TestCard_Suit(t *testing.T) {
	assert.Equal(t, "Schellen", SchellenOber.Suit().String(), "Card names should be correct")
}

func TestCard_CardValue(t *testing.T) {
	assert.Equal(t, "Ober", SchellenOber.CardValue().String(), "Card names should be correct")
}