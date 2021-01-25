package main

import (
	"bytes"
	"fmt"
	"math"
)

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

示例 1：
输入：strs = ["flower","flow","flight"]
输出："fl"

示例 2：
输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。

提示：

0 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
通过次数435,171提交次数1,112,831

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	strs := []string{"flowe", "flower", "flower", "flower"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	minSize := math.MaxInt64
	for _, str := range strs {
		if minSize > len(str) {
			minSize = len(str)
		}
	}

	var buf bytes.Buffer
	for i := 0; i < minSize; i++ {
		next := check(strs, i)
		if next != 0 {
			buf.WriteByte(next)
		} else {
			break
		}
	}
	return buf.String()
}

func check(strs []string, idx int) byte {
	m := map[byte]int{}
	var now byte
	for _, str := range strs {
		if len(str) == 0 || len(str) < idx+1 {
			return 0
		}
		now = str[idx]
		m[now]++
	}
	if now != 0 {
		if m[now] != len(strs) {
			return 0
		} else {
			return now
		}
	} else {
		return 0
	}
}
