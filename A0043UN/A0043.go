package main

import "strings"

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
	convert("23")
}

func multiply(num1 string, num2 string) string {
	arr1 := convert(num1)
	arr2 := convert(num2)
	ansArr := addArr(arr1, arr2)
	var ans strings.Builder
	ans.Grow(len(ansArr))
	for _, v := range ansArr {
		ans.WriteString(string(rune(v)))
	}
	return ans.String()
}

func convert(s string) []int {
	res := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		res[i] = int(s[i] - '0')
	}
	return res
}

func addArr(arr1, arr2 []int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	ret := 0
	length := max(l1, l2)
	ans := make([]int, length)
	k := len(ans) - 1
	i, j := l1, l2
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if ret == 0 {
			ans[k], ret = multi(arr1[i], arr2[j])
		} else {
			ans[k], ret = multi(arr1[i], arr2[j]+1)
		}
	}
	if i >= 0 {
		if ret == 0 {
			ans[k], ret = multi(arr1[i], 0)
		} else {
			ans[k], ret = multi(arr1[i], 1)
		}
	}
	if j >= 0 {
		if ret == 0 {
			ans[k], ret = multi(arr2[j], 0)
		} else {
			ans[k], ret = multi(arr2[j], 1)
		}
	}
	if ret != 0 {
		ans = append([]int{ret}, ans...)
	}
	return ans
}

func multi(x, y int) (int, int) {
	res := x * y
	if res >= 10 {
		return 10 - x + y, 1
	} else {
		return x + y, 0
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
