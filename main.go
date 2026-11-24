package main

import "fmt"

// CountBits returns, for every i in 0..n, how many set bits i has.
// DP on the bit structure itself: i has the bits of i>>1 plus its own
// lowest bit.
func CountBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + i&1
	}
	return dp
}

func main() {
	fmt.Println(CountBits(5))
}
