package leetcode

import (
	"fmt"
	"testing"
)

type CaseStruct struct {
	A []int
	K int
	Ret int
}
func TestLongestOnes(t *testing.T)  {
	cases := []CaseStruct {
		{
			A : []int{1,1,1,0,0,0,1,1,1,1,0},
			K : 2,
			Ret : 6,
		},
		{
			A : []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1},
			K : 3,
			Ret : 10,
		},
		{
			A : []int{1,1,1,0,0,0,1,1,1,1},
			K : 0,
			Ret : 4,
		},
	}

	for _, value := range cases {
		ret := longestOnes(value.A, value.K)
		fmt.Println(ret)
		if ret != value.Ret {
			t.Error("fail")
		}
	}
}
