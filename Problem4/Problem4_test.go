package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostAppearItem(t *testing.T) {

	t.Run("case1: ", func(t *testing.T) {
		assert.NotEqual(t, string("golang->1 ruby->2 js->4"), MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}), "Hasil output tidak sesuai")
	})
	t.Run("case3: ", func(t *testing.T) {
		assert.NotEqual(t, string("A->1 D->3 C->4 B->2"), MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}), "Hasil output tidak sesuai")
	})
}
