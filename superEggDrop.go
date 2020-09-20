package main

import "fmt"

func superEggDrop(K int, N int) int {
	dp := make([]int, N+1)
	for i:=0; i<cap(dp); i++ {
		dp[i] = i
	}
	for k:=2; k<=K; k++ {
		dp2 := make([]int, N+1)
		x := 1
		for n:=1; n<=N; n++ {
			for x<n && maxFunc(dp[x-1], dp2[n-x]) > maxFunc(dp[x], dp2[n-x-1]) {
				x ++
			}
			dp2[n] = 1 + maxFunc(dp[x-1], dp2[n-x])
		}
		dp = dp2
	}
	return dp[N]
}

func maxFunc(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(superEggDrop(2,9))
}
