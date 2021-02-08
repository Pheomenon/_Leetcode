package main

/*
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true
限制：

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。
通过次数54,761提交次数75,862

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-unique-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := "abc"
	isUnique(s)
}

func isUnique(astr string) bool {
	if len(astr) == 0 {
		return true
	}
	var mark uint32
	mark = 1
	m := astr[0] - 'a'
	mark <<= m
	for i := 1; i < len(astr); i++ {
		var tmp uint32
		tmp = 1
		r := astr[i] - 'a'
		tmp <<= r
		mt := mark
		if mt == tmp|mark {
			return false
		} else {
			mark |= tmp
		}
	}
	return true
}
