package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOver(t *testing.T) {
	sum := 0
	for v := range Over(1, 2, 3) {
		sum += v
	}
	assert.Equal(t, 6, sum)
}
