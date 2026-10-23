package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullDeck(t *testing.T) {
	d := Deck{}
	d.fill()
	assert.Equal(t, 54, len(d.cards))
}