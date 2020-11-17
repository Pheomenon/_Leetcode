/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"

提示：
1 <= s.length, t.length <= 105
s 和 t 由英文字母组成

进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
通过次数88,790提交次数224,083

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	s := "aab"
	t := "aab"
	fmt.Println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	window, need := map[byte]int{}, map[rune]int{}
	for _, char := range t {
		need[char]++
	}
	left, right, valid := 0, 0, 0
	//记录最小覆盖子串的起始索引和长度
	start, length := 0, math.MaxInt64
	for right < len(s) {
		//c 是将移入窗口的字符
		c := s[right]
		//右移窗口
		right++
		v, ok := need[rune(c)]
		if ok {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		//DEBUG
		//fmt.Printf("window: [%d, %d]\n", left, right)

		for valid == len(need) {
			//更新最小覆盖子串
			if right-left < length {
				start = left
				length = right - left
			}
			//d是移出窗口的字符
			d := s[left]
			left++
			v, ok := need[rune(d)]
			if ok {
				if window[d] == v {
					valid--
				}
				window[d]--
			}
		}

		/**
		for window needs shrink {
			//d 是将移出窗口的字符
			d := s[left]
			//左移窗口
			left++
		}*/
	}
	if length == math.MaxInt64 {
		return ""
	}
	return s[start : start+length]
}
