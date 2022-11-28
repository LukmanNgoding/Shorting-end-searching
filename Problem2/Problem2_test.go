package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBuyProduct(t *testing.T) {
	t.Run("case1: 50000, []int{25000, 25000, 10000, 14000}", func(t *testing.T) {
		assert.Equal(t, int(3), MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}), "Hasil output tidak sesuai")
	})
	t.Run("case2: 30000, []int{15000, 10000, 12000, 5000, 3000}", func(t *testing.T) {
		assert.NotEqual(t, int(3), MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}), "Hasil output tidak sesuai")
	})

}
