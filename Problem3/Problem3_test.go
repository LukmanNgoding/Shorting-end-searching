package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayingDomino(t *testing.T) {
	t.Run("case1: [][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}", func(t *testing.T) {
		assert.Equal(t, []int{3, 4}, PlayingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}), "Hasil output tidak sesuai")
	})
	t.Run("case1: [][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}", func(t *testing.T) {
		assert.NotEqual(t, []int{2, 1}, PlayingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}), "Hasil output tidak sesuai")
	})

}
