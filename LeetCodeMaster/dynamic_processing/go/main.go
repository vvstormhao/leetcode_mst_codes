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

	dp[1] = 1
	dp[2] = 1
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

func max(a, b int) int {
	if a > b {
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

/*
* 整数拆分
给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。
示例 1: 输入: 2 输出: 1
\解释: 2 = 1 + 1, 1 × 1 = 1。
示例 2: 输入: 10 输出: 36 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。 说明: 你可以假设 n 不小于 2 且不大于 58。
*/

/*
解题思路
* dp[i]含义：整数拆分后的道的最大乘积
* 递推公式: dp[i] = max(i * [j - i], dp[i - j] * j)
* dp初始化：dp[0]和dp[1]没有意义，从dp[2]开始初始化，dp[2] = 1
*/

func IntegerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(i*(i-j), dp[i-j]*j))
		}
	}

	return dp[n]
}

/*
*不同的二叉树拆分
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

*解题思路
1.dp[i]含义：从0...i 为节点组成的二叉搜索树种类
2.递推公式： dp[i] += dp[j - 1] * dp[i - j];
3.初始化: dp[0] = 0
*/
func numTrees(n int) int {
	dp := make([]int, n + 1)
	dp[0] = 0

	for i := 1;i <=n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j - 1] * dp[i - j]
		}
	}

	return dp[n]
}



// 01背包问题
/*
* 标准的背包问题
在下面的讲解中，我举一个例子：
背包最大重量为4。
物品为：
重量	价值
物品0	1	15
物品1	3	20
物品2	4	30
问背包能背的物品最大价值是多少？

*解题思路
* dp[i][j]含义， 从0...i ，当背包的容量为j时，所能容纳的最大价值
* 递推公式：dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]] + value[i])
*/
func BasicPackage01UsingIJ(bagWeight int) {
	weight := []int{1,3,4}
	value := []int{15, 20, 30}

	dp := make([][]int, bagWeight + 1)
	for i := 0;i < len(weight); i++ {
		v := make([]int, bagWeight + 1)
		dp = append(dp, v)
	}

	for j := weight[0]; j <= bagWeight; j++ {
		dp[0][j] = value[0]
	}

	for i := 0;i <= len(weight);i++ {
		dp[i][0] = 0
	}

	for i := 1;i <= len(weight);i++ {
		for j := 1;j <= bagWeight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i - 1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]] + value[i])
			}
		}
	}

	fmt.Printf("%v\n", dp)
}

/*
* 01基础背包问题，通过滚动数组实现
* 解题思路
* dp[j]含义：当背包容量为j时，背包所装下物品的最大价值
* 递推公式：dp[j] = max(dp[j], dp[j-weight[i]] + value[i])
* 初始化：dp[0] = 0,共2层循环，外层循环为物品，内层循环为背包质量。内层循环背包质量从大到小遍历，目的是为了保证每个物品只被计算一次。
*/

func BasicPackage01UsingJ(bagWeight int) {
	weight := []int{1,3,4}
	value := []int{15, 20, 30}

	dp := make([]int, bagWeight + 1)
	for j := 0; j <= bagWeight; j++ {
		dp[j] = 0
	}

	for i := 0; i < len(weight); i++ {
		for j := bagWeight; j >= weight[i]; j-- {
				dp[j] = max(dp[j], dp[j-weight[i]] + value[i])
				fmt.Printf("weight[%d] = %d, dp[%d] = %d\n", i, weight[i], j, dp[j])
		}
	}

	fmt.Printf("final %v\n", dp)
}

/*
分割等和子集
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
注意: 每个数组中的元素不会超过 100 数组的大小不会超过 200
示例 1: 输入: [1, 5, 11, 5] 输出: true 解释: 数组可以分割成 [1, 5, 5] 和 [11].
示例 2: 输入: [1, 2, 3, 5] 输出: false 解释: 数组不能分割成两个元素和相等的子集.

* 解题思路
先将数组中的元素排序，然后计算总和，总和/2即为背包的容量。物品重量和价值都是一样的。
*/

func CanPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}

	var target int
	if sum % 2 == 1 {
		return false
	}

	target = sum / 2

	dp := make([]int, target + 1)
	dp[0] = 0

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i];j-- {
			dp[j] = max(dp[j], dp[j-nums[i]] + nums[i])
		}
	}

	if dp[target] == target {
		return true
	}

	return false
}

func main() {
	//result := fib(10)
	//fmt.Printf("result %d\n", result)

	BasicPackage01UsingJ(7)
}
