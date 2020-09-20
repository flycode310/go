package leetcode

import "fmt"

func maxProfit(prices []int) int {
	lowestPrice := int(0x7fffffff)
	maxProfit := 0
	profit := 0
	for _, v := range prices {
		if v < lowestPrice {
			lowestPrice = v
		}
		if v > lowestPrice {
			profit = v - lowestPrice
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func MaxProfit() {
	prices := []int{7,1,5,3,6,4}
	result := maxProfit(prices)
	fmt.Println(result)
}
