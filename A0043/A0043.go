package main

import (
	"fmt"
	"strings"
)

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，
返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。
示例 1:
输入: num1 = "2", num2 = "3"
输出: "6"

示例 2:
输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于 110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
通过次数123,480提交次数276,222

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/multiply-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	fmt.Println(multiply("5423", "564"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ansArr := addArr(num1, num2)
	var ans strings.Builder
	ans.Grow(len(ansArr))
	for _, v := range ansArr {
		ans.WriteByte(v + '0')
	}
	return ans.String()
}

func addArr(arr1, arr2 string) []uint8 {
	ans := make([]uint8, len(arr2)+len(arr1))
	for i := len(arr1) - 1; i >= 0; i-- {
		for j := len(arr2) - 1; j >= 0; j-- {
			r := (arr1[i] - '0') * (arr2[j] - '0')
			r += ans[i+j+1]
			ans[i+j+1] = r % 10
			ans[i+j] += r / 10
		}
	}
	for _, an := range ans {
		if an == 0 {
			ans = ans[1:]
		} else {
			break
		}
	}
	return ans
}
