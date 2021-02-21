package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"
通过次数79,902提交次数147,944

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	str := "abc3[cd]xyz"
	fmt.Println(decodeString(str))
}

func decodeString(s string) string {
	stackStr := []rune{}
	stackNum := []int{}
	base := 0
	//var strContent strings.Builder
	for _, char := range s {
		if unicode.IsDigit(char) {
			num := char - '0'
			base = base*10 + int(num)
		} else if unicode.IsLetter(char) {
			stackStr = append(stackStr, char)
		} else if char == '[' {
			stackNum = append(stackNum, base)
			stackStr = append(stackStr, char)
			base = 0
		} else if char == ']' {
			times := stackNum[len(stackNum)-1]
			stackNum = stackNum[:len(stackNum)-1]
			content := make([]rune, 0)
			for i := len(stackStr) - 1; i > 0; i-- {
				if stackStr[i] != '[' {
					content = append(content, stackStr[i])
					stackStr = stackStr[:len(stackStr)-1]
				} else {
					stackStr = stackStr[:len(stackStr)-1]
					break
				}
			}
			reverse(content)
			rep := strings.Repeat(string(content), times)

			for _, i2 := range rep {
				stackStr = append(stackStr, i2)
			}
		}
	}
	for i := 0; i < len(stackStr); i++ {
		if stackStr[i] == '[' {
			stackStr = stackStr[1:]
		} else {
			break
		}
	}
	return string(stackStr)
}

func reverse(content []rune) []rune {
	for i, j := 0, len(content)-1; i < j; i, j = i+1, j-1 {
		content[i], content[j] = content[j], content[i]
	}
	return content
}
