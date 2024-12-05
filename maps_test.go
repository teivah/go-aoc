package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapGet(t *testing.T) {
	m := map[int]map[int]int{
		1: {
			10: 100,
		},
	}
	assert.Equal(t, map[int]int{10: 100}, InnerMapGet(m, 1))
	assert.Equal(t, map[int]int{}, InnerMapGet(m, 2))
	assert.Equal(t, map[int]int{}, m[2])
	assert.Equal(t, map[int]int{}, InnerMapGet(m, 3, 10))
}
