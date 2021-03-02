package leetcode

import (
	"fmt"
	"testing"
)

func TestOffer29(t *testing.T) {
	matrix := [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	ant := spiralOrder(matrix)
	fmt.Println(ant)
}
