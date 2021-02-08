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
	if len(digits) == 0 {
		return []string{}
	}
	ans := []string{}
	alphbet := map[uint8]string{
		2: "abc",
		3: "def",
		4: "ghi",
		5: "jkl",
		6: "mno",
		7: "pqrs",
		8: "tuv",
		9: "wxyz",
	}

	var backtrack func(track, digits string)
	backtrack = func(track, digits string) {
		if len(digits) == 0 {
			ans = append(ans, track)
			return
		}

		for _, v := range alphbet[digits[0]-'0'] {
			track += string(v)
			backtrack(track, digits[1:])
			track = track[:len(track)-1]
		}
	}
	backtrack("", digits)

	return ans
}
