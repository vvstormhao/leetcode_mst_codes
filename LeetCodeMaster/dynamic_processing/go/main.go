package main

import "fmt"

// 动态规划

/*
* 斐波那契数列
* dp[i] 代表斐波那契数列中第i个值的大小
* 递推公式: dp[i] = dp[i - 2] + dp[i - 1]
* dp初始化 dp[0] = 0 dp[1] = 1
* 遍历方向：从前往后
 */

func fib(n int) int {
	dp := make([]int, n, n)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func main() {
	result := fib(10)
	fmt.Printf("result %d\n", result)
}
