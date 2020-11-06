/**
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。
注意:

给定数字的范围是 [0, 10^8]
通过次数9,578提交次数22,836

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-swap
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import (
	"fmt"
)

func main() {
	num := 92321321 // 92321
	fmt.Println(maximumSwap(num))
}

func maximumSwap(num int) int {
	if num == 0 {
		return 0
	}
	//按位把原数字录入到数组中
	res := make([]int, 0, 8)
	for num != 0 {
		res = append(res, num%10)
		num /= 10
	}
	//从最高位开始找到一个比当前位大而且是最低位的数字。
	//比如12324323当指针在最高位时不和离它最近3交换而是和最低位3交换这样保证了交换后的数字是最大的
	for i := len(res) - 1; i >= 0; i-- {
		if findMaxAndSwap(res, i) {
			break
		}
	}
	//还原数组中的数字
	result := res[len(res)-1]
	for i := len(res) - 2; i >= 0; i-- {
		result *= 10
		result += res[i]
	}
	return result
}

func findMaxAndSwap(res []int, swapIndex int) bool {
	//swapIndex指的是此时将被比较的最高位数字
	if swapIndex == 0 {
		return true
	}
	max := res[swapIndex]
	exIndex := swapIndex
	for i := swapIndex; i >= 0; i-- {
		if max <= res[i] {
			max = res[i]
			exIndex = i
		}
	}
	//成功找到最低位的最大数，返回true
	if res[swapIndex] != res[exIndex] {
		tmp := res[swapIndex]
		res[swapIndex] = max
		res[exIndex] = tmp
		return true
	} else {
		return false
	}
}
