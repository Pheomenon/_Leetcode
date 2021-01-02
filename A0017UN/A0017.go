package main

import "fmt"

/*
给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

通过次数206,244提交次数370,778

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	digit := "23"
	fmt.Println(letterCombinations(digit))
}

func letterCombinations(digits string) []string {
	var alphabet = []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	var track string
	ans := make([]string, 0)

	var backtrack func(Kstart int, track string)
	backtrack = func(Kstart int, track string) {

		if len(track) == len(digits) {
			return
		}

		for i := 0; i < len(digits); i++ {
			for j := i; j < len(alphabet[digits[i]]); j++ {
				track += alphabet[digits[j]]
				for k := Kstart; k < len(alphabet[digits[k]]); k++ {
					track += alphabet[digits[k]]
					ans = append(ans, track)
					track = track[:len(track)-len(alphabet[digits[k]])]
				}
				backtrack(Kstart+1, track)
			}
		}
	}
	backtrack(0, track)
	return ans
}
