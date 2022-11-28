package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinAndMax(t *testing.T) {
	t.Run("case1: 1", func(t *testing.T) {
		assert.Equal(t, "min: -2 index: 3 max: 8 index: 5", FindMinAndMax([]int{5, 7, 4, -2, -1, 8}), "Hasil output tidak sesuai")
	})
	t.Run("case2: 2", func(t *testing.T) {
		assert.NotEqual(t, "min: -5 index: 1 max: 22 index: 4", FindMinAndMax([]int{2, -5, -4, 22, 7, 7}), "Hasil output tidak sesuai")
	})

}
