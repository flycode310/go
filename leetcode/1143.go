package leetcode

import (
	"fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	table := make([][]int, len1+1)
	for i:=0; i<=len1; i++ {
		table[i] = make([]int, len2+1)
	}
	table[0][0] = 0
	for i:=1; i<=len1; i++ {
		table[i][0] = 0
	}
	for j:=1; j<=len2; j++ {
		table[0][j] = 0
	}

	for i:=1; i<=len1; i++ {
		for j:=1; j<=len2; j++ {
			if text1[i-1] == text2[j-1] {
				table[i][j] = table[i-1][j-1] + 1
			} else {
				if (table[i][j-1] > table[i-1][j]) {
					table[i][j] = table[i][j-1]
				} else {
					table[i][j] = table[i-1][j]
				}
			}
		}
	}

	return table[len1][len2]
}

func LongestCommonSubsequence() {
	text1 := "abc"
	text2 := "def"
	ret := longestCommonSubsequence(text1, text2)
	fmt.Println(ret)
}
