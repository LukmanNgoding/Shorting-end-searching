package main

import (
	"fmt"
)

func PlayingDomino(cards [][]int, deck []int) []int {

	var hasil []int
	for i := 0; i < len(cards); i++ {
		if deck[1] == cards[i][0] {
			hasil = append(hasil, cards[i][0])
			hasil = append(hasil, cards[i][1])
			break
		}

	}
	return hasil
}

func main() {
	fmt.Println(PlayingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	// [3, 4]
	fmt.Println(PlayingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	// [6 5]
	fmt.Println(PlayingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
