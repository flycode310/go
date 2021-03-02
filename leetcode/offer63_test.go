package leetcode

import (
	"fmt"
	"testing"
)

func TestOffer63(t *testing.T) {
	prices := []int{7,6,4,3,1}
	ant := maxProfit63(prices)
	fmt.Println(ant)
}
