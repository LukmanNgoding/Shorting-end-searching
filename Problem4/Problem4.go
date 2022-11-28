package main

import (
	"fmt"
	"sort"
	"strconv"
)

func MostAppearItem(item []string) string {
	sort.Strings(item)
	var temp []string
	var temp1 string = item[0]
	var in int

	temp = append(temp, item[0])

	for i := 1; i < len(item); i++ {
		if temp1 < item[i] {
			temp = append(temp, item[i])
			temp1 = item[i]
		}
	}
	dict := make(map[string]int)
	for _, nun := range item {
		dict[nun] = dict[nun] + 1
	}
	v := make([]int, 0, len(dict))
	for _, value := range dict {
		v = append(v, value)
	}

	for i := 0; i < len(v)-1; i++ {
		if v[i] < v[i+1] {
			in = v[i]
			v[i] = v[i+1]
			v[i+1] = in
			temp1 = temp[i]
			temp[i] = temp[i+1]
			temp[i+1] = temp1
		}
	}
	var hasil string
	for i := len(v) - 1; i >= 0; i-- {
		hasil += temp[i] + "->"
		hasil += strconv.Itoa(v[i]) + " "
	}
	return hasil
}

func main() {

	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
