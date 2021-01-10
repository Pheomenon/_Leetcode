package main

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例：

输入："Let's take LeetCode contest"
输出："s'teL ekat edoCteeL tsetnoc"

提示：

在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。
通过次数107,033提交次数145,219

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	var ans string
	pos := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ans += reverse(s[pos:i])
			ans += " "
			pos = i + 1
		}
		if i == len(s)-1 {
			i++
			ans += reverse(s[pos:i])
		}
	}
	return ans
}

func reverse(s string) string {
	builder := &strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		builder.WriteByte(s[i])
	}
	return builder.String()
}
