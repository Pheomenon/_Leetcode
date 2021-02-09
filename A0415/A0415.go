package main

import (
	"fmt"
	"strings"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：

num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库，也不能直接将输入的字符串转换为整数形式
通过次数91,109提交次数175,654

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	num1 := "9"
	num2 := "99"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	} else if num2 == "0" {
		return num1
	}
	num1L := len(num1) - 1
	num2L := len(num2) - 1
	ans := make([]uint8, max(num1L, num2L)+3)
	ansI := len(ans) - 1

	for num1L >= 0 || num2L >= 0 {
		if num1L >= 0 && num2L >= 0 {
			r := num1[num1L] - '0' + num2[num2L] - '0'
			r += ans[ansI]
			ans[ansI] = r % 10
			ans[ansI-1] = r / 10
			num1L--
			num2L--
			ansI--
		} else if num1L < 0 && num2L >= 0 {
			r := num2[num2L] - '0'
			r += ans[ansI]
			ans[ansI] = r % 10
			ans[ansI-1] = r / 10
			num2L--
			ansI--
		} else if num2L < 0 && num1L >= 0 {
			r := num1[num1L] - '0'
			r += ans[ansI]
			ans[ansI] = r % 10
			ans[ansI-1] = r / 10
			num1L--
			ansI--
		}
	}

	for _, an := range ans {
		if an == 0 {
			ans = ans[1:]
		} else {
			break
		}
	}

	var ansS strings.Builder

	for i := 0; i < len(ans); i++ {
		ansS.WriteByte(ans[i] + '0')
	}
	return ansS.String()
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
