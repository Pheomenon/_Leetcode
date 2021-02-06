package main

import "fmt"

/*
给你一个仅由大写英文字母组成的字符串，你可以将任意位置上的字符替换成另外的字符，
总共可最多替换 k 次。在执行上述操作后，找到包含重复字母的最长子串的长度。
注意：字符串长度 和 k 不会超过 10^4。

示例 1：
输入：s = "ABAB", k = 2
输出：4
解释：用两个'A'替换为两个'B',反之亦然。

示例 2：
输入：s = "AABABBA", k = 1
输出：4
解释：
将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
子串 "BBBB" 有最长重复字母, 答案为 4。
通过次数17,691提交次数35,197

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-repeating-character-replacement
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := "ABBB"
	k := 2
	fmt.Println(characterReplacement(s, k))
}

func characterReplacement(s string, k int) int {
	left, mid, right := 0, 1, 2
	ans := 0
	ky := k
	for i := 0; i < len(s)-1; i++ {
		k = ky
		mid = i
		left = mid
		right = mid
		for left > -1 || right < len(s) || k != 0 {
			if left > -1 {
				if s[left] == s[mid] {
					left--
				} else if s[left] != s[mid] && k > 0 {
					left--
					k--
				}
			}
			if right < len(s) {
				if s[right] == s[mid] {
					right++
				} else if s[right] != s[mid] && k > 0 {
					right++
					k--
				}
			}
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
