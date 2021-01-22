/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
通过次数415,998提交次数1,288,360
*/
package main

import (
	"fmt"
)

func main() {
	s := "aa"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	L1, R1, L2, R2 := 0, 0, 0, 0
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		L1, R1 = check(s, i, i)
		L2, R2 = check(s, i, i+1)
		if R1-L1 > end-start {
			start, end = L1, R1
		}
		if R2-L2 > end-start {
			start, end = L2, R2
		}
	}
	return s[start : end+1]
}

func check(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
