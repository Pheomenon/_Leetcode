package main

import (
	"fmt"
)

/*
如果一个十进制数字不含任何前导零，且每一位上的数字不是 0 就是 1 ，那么该数字就是一个 十-二进制数 。例如，101 和 1100 都是 十-二进制数，而 112 和 3001 不是。

给你一个表示十进制整数的字符串 n ，返回和为 n 的 十-二进制数 的最少数目。

示例 1：

输入：n = "32"
输出：3
解释：10 + 11 + 11 = 32
示例 2：

输入：n = "82734"
输出：8
示例 3：

输入：n = "27346209830709182346"
输出：9

提示：

1 <= n.length <= 105
n 仅由数字组成
n 不含任何前导零并总是表示正整数
通过次数5,841提交次数6,669

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	n := "27346209830709182346"
	fmt.Println(minPartitions(n))
}

func minPartitions(n string) int {
	var max byte
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
		if max == 9 {
			return 9
		}
	}
	return int(max - '0')
}
