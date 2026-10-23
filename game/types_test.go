package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberOfFaces(t *testing.T) {
	faces := LoadFaces()
	assert.Equal(t, 16, len(faces))
}