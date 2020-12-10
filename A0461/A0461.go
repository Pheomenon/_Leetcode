package main

import (
	"fmt"
)

/*
两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
给出两个整数 x 和 y，计算它们之间的汉明距离。

注意：
0 ≤ x, y < 231.

示例:
输入: x = 1, y = 4
输出: 2

解释:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑

上面的箭头指出了对应二进制位不同的位置。
通过次数85,054提交次数108,659

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(hammingDistance(3, 1))
}

func hammingDistance(x int, y int) int {
	xb := fmt.Sprintf("%b", x)
	yb := fmt.Sprintf("%b", y)

	length := max(len(xb), len(yb))
	res := max(length-len(xb), length-len(yb))
	if len(xb) != length {
		for i := 0; i < res; i++ {
			xb = "0" + xb
		}
	}
	if len(yb) != length {
		for i := 0; i < res; i++ {
			yb = "0" + yb
		}
	}

	cnt := 0
	for i := 0; i < len(xb); i++ {
		if xb[i] != yb[i] {
			cnt++
		}
	}
	return cnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
