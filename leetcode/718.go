package leetcode

import "fmt"

func findLength(A []int, B []int) int {
	dp := [1001][1001]int{}
	lenA := len(A)
	lenB := len(B)
	if lenA == 0 || lenB == 0 {
		return 0
	}

	maxLen := 0
	for i:=1; i<=lenA; i++ {
		for j:=1; j<=lenB; j++ {
			indexA := i - 1
			indexB := j - 1
			if A[indexA] == B[indexB] {
				if dp[i-1][j-1] > 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
				if maxLen < dp[i][j] {
					maxLen = dp[i][j]
				}
			}
		}
	}

	return maxLen
}

func FindLength() {
	A := []int{1,2,3,2,1}
	B := []int{3,2,1,4,7}
	ret := findLength(A, B)
	fmt.Println(ret)
}
