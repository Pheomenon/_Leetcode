/**
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
如果数组元素个数小于 2，则返回 0。

示例 1:
输入: [3,6,9,1]
输出: 3
解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。

示例 2:
输入: [10]
输出: 0
解释: 数组元素个数小于 2，因此返回 0。

说明:
你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
通过次数20,317提交次数36,823

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-gap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

//桶排序
import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 6, 9, 1}
	fmt.Println(maximumGap(nums))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	matrix := [10000][10000]int{{}}
	for _, v := range nums {
		matrix[v][v] = v
	}
	max := math.MinInt64
	res := make([]int, 0, 1024)
	for i := 0; i < len(matrix); i++ {
		if matrix[i][i] != 0 {
			res = append(res, matrix[i][i])
		}
	}
	ret := 0
	for i := len(res) - 1; i > 1; i-- {
		ret = res[i] - res[i-1]
		if ret > max {
			max = ret
		}
	}
	return max
}
