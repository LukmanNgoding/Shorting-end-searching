package main

import (
	"fmt"
	"sort"
)

func MaximumBuyProduct(money int, productPrice []int) int {

	var count int
	sort.Ints(productPrice)
	for i := 0; i < len(productPrice) && money >= 0; i++ {
		if (money - int(productPrice[i])) >= 0 {
			money = money - int(productPrice[i])
			count++
		}
	}
	return count

}

func main() {

	fmt.Println(MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))      //3
	fmt.Println(MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})) //4
	fmt.Println(MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000}))   //4
	fmt.Println(MaximumBuyProduct(4000, []int{75000, 3000, 25000, 2000}))         //1
	fmt.Println(MaximumBuyProduct(0, []int{10000, 30000}))                        //0
}