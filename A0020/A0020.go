package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
通过次数488,998提交次数1,124,352

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type stack struct {
	size    int
	content []uint8
}

func main() {
	s := "]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := &stack{
		content: make([]uint8, len(s)),
		size:    -1,
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' || s[i] == ' ' {
			stack.size++
			stack.content[stack.size] = s[i]
			continue
		}
		if stack.size == -1 {
			return false
		}
		if s[i] == ')' {
			if stack.content[stack.size] == uint8('(') {
				stack.size--
			} else {
				return false
			}
			continue
		}
		if s[i] == ']' {
			if stack.content[stack.size] == uint8('[') {
				stack.size--
			} else {
				return false
			}
			continue
		}
		if s[i] == '}' {
			if stack.content[stack.size] == uint8('{') {
				stack.size--
			} else {
				return false
			}
			continue
		}
		if s[i] == ' ' {
			if stack.content[stack.size] == uint8(' ') {
				stack.size--
			} else {
				return false
			}
			continue
		}
	}
	if stack.size != -1 {
		return false
	}
	return true
}
