/**
给定一个字符串 s 和一个非空字符串 p，
找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。

说明：
字母异位词指字母相同，但排列不同的字符串。
不考虑答案输出的顺序。

示例 1:
输入:
s: "cbaebabacd" p: "abc"
输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。

示例 2:
输入:
s: "abab" p: "ab"
输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
通过次数44,442提交次数93,773

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-all-anagrams-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
	s := "eidbaooo"
	p := "ab"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s string, p string) []int {
	window, need := map[rune]int{}, map[rune]int{}
	left, right, valid := 0, 0, 0
	result := make([]int, 0, 128)
	for _, s := range p {
		need[s]++
	}
	for _, c := range s {
		right++
		v, ok := need[c]
		if ok {
			window[c]++
			if window[c] == v {
				valid++
			}
		}

		//DEBUG
		//fmt.Printf("window: [%d, %d]\n", left, right)

		for right-left >= len(p) {
			if valid == len(need) {
				result = append(result, left)
			}
			d := s[left]
			left++
			v, ok := need[rune(d)]
			if ok {
				if window[rune(d)] == v {
					valid--
				}
				window[rune(d)]--
			}
		}
	}
	return result
}
