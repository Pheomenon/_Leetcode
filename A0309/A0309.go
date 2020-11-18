/**
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，
你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

示例:
输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]
通过次数63,451提交次数110,919

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	//dp_pre_0指两天以前
	dp_pre_0, dp_i_0, dp_i_1 := 0, 0, math.MinInt32
	for i := 0; i < len(prices); i++ {
		tmp := dp_i_0
		dp_i_0 = max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = max(dp_i_1, dp_pre_0-prices[i])
		dp_pre_0 = tmp
	}
	return dp_i_0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
