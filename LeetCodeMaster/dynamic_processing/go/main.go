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

/*
* 爬楼梯
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1 阶 + 1 阶
2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1 阶 + 1 阶 + 1 阶
1 阶 + 2 阶
2 阶 + 1 阶
*/

/*
*解题思路
*1.确定dp数组：dp[i]表示爬到第i层的方法
*2.递推公式：dp[i] = dp[i-2] + dp[i - 1]
*3.初始化：
   dp[0] = 1
   dp[1] = 1
   dp[2] = 2
*4.确定循环遍历的顺序，从前往后。
*5.推导dp数组：
*/
func getMaxClimbWay(n int) int {
	dp := make([]int, n+1, n+1)
	if n <= 1 {
		return n
	}

	dp[1] == 1
	dp[2] == 1
	for i := 2; i <= n;i++ {
		dp[i] = dp[i - 2] + dp[i - 1]
	}

	return dp[n]
}

/*
*最小花费爬楼梯
数组的每个下标作为一个阶梯，第 i 个阶梯对应着一个非负数的体力花费值 cost[i]（下标从 0 开始）。

每当你爬上一个阶梯你都要花费对应的体力值，一旦支付了相应的体力值，你就可以选择向上爬一个阶梯或者爬两个阶梯。

请你找出达到楼层顶部的最低花费。在开始时，你可以选择从下标为 0 或 1 的元素作为初始阶梯。

示例 1：

输入：cost = [10, 15, 20] 输出：15 解释：最低花费是从 cost[1] 开始，然后走两步即可到阶梯顶，一共花费 15 。  示例 2：

输入：cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1] 输出：6 解释：最低花费方式是从 cost[0] 开始，逐个经过那些 1 ，跳过 cost[3] ，一共花费 6 。
*/

/*
* 解题思路
* 1.确定dp数组：di[i]表示爬到i所需最小体力
* 2.递推公式:dp[i] = min(dp[i - 1],dp[i - 2]) + cost[i]
* 3.初始化：dp[0] = cost[0],dp[1] = cost[1]
* 4.遍历顺序：从前往后
*/
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func climbStiresCosts(costs []int, n int) int {
	dp := make([]int, n + 1)
	dp[0] = costs[0]
	dp[1] = costs[1]
	for i := 2;i<=n;i++ {
		dp[i] = min(dp[i - 2], dp[i - 1]) + costs[i]
	}

	return dp[n]
}

/*
* 不同路径
一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
问总共有多少条不同的路径？
示例1：
输入：m = 3, n = 7
输出：28

示例 2：
输入：m = 2, n = 3
输出：3
*/

/*
* 解题思路
* dp[i][j]:位于i*j位置的路径总数
* 递推公式:dp[i][j] = dp[i-1][j] + dp[i][j-1], 因为只能从左边或上边这两个路径过来
* 初始化：dp[0][j] = 1, dp[i][0] = 1
* 遍历顺序：从前往后，从上往下
*/

func diffWays(route [][]int, m,n int) int {
	dp := make([][]int, m)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int,n)
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]

}

func main() {
	result := fib(10)
	fmt.Printf("result %d\n", result)
}
