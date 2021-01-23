package main

/*
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

提示：
s.length <= 40000
注意：本题与主站 3 题相同：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

通过次数68,558提交次数149,407

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := "anviaj"
	lengthOfLongestSubstring(s)
}

func lengthOfLongestSubstring(s string) int {
	right := -1
	window := map[byte]bool{}
	ans := 0
	//实际上就是维护一个窗口存储每个子串，然后从每个字符出发看能延伸多长
	for i := 0; i < len(s); i++ {
		if i != 0 {
			delete(window, s[i-1])
		}
		for right+1 < len(s) && !window[s[right+1]] {
			window[s[right]] = true
			right++
		}
		ans = max(ans, right-i)
	}
	return ans
}

/*每次循环新建Map的版本：*/
//func lengthOfLongestSubstring(s string) int {
//	 right := 1
//	ans := 0
//	//实际上就是维护一个窗口存储每个子串，然后从每个字符出发看能延伸多长
//	for i := 0; i < len(s); i++ {
//		window := map[byte]bool{}
//		window[s[i]] = true
//		right = i+1
//		for right < len(s) && !window[s[right]] {
//			window[s[right]] = true
//			right++
//		}
//
//		ans = max(ans, right-i)
//	}
//	return ans
//}
//
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
