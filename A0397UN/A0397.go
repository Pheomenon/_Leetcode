/*
给定一个正整数 n ，你可以做如下操作：
如果 n 是偶数，则用 n / 2替换 n 。
如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
n 变为 1 所需的最小替换次数是多少？

示例 1：
输入：n = 8
输出：3
解释：8 -> 4 -> 2 -> 1

示例 2：
输入：n = 7
输出：4
解释：7 -> 8 -> 4 -> 2 -> 1
或 7 -> 6 -> 3 -> 2 -> 1

示例 3：
输入：n = 4
输出：2

提示：

1 <= n <= 231 - 1
通过次数9,899提交次数26,955

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	s := 65535
	fmt.Println(integerReplacement(s))
}

func integerReplacement(n int) int {
	return min(recursionAdd(n), recursionSub(n))
}

func recursionAdd(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n + 1
		}
		count++
	}
	return count
}

func recursionSub(n int) int {
	count := 0
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n - 1
		}
		count++
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
