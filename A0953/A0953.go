/**
某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。

给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，
返回 true；否则，返回 false。

示例 1：

输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
输出：true
解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。
示例 2：

输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
输出：false
解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。
示例 3：

输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
输出：false
解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，因为 'l' > '∅'，其中 '∅' 是空白字符，定义为比任何其他字符都小（更多信息）。

提示：

1 <= words.length <= 100
1 <= words[i].length <= 20
order.length == 26
在 words[i] 和 order 中的所有字符都是英文小写字母。
通过次数8,695提交次数15,386

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/verifying-an-alien-dictionary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	words := []string{"word", "world", "row"}
	order := "worldabcefghijkmnpqstuvxyz"
	fmt.Println(isAlienSorted(words, order))
}

//遍历words数组，取出元素的第一个字符，然后根据order给res数组赋值，
//res数组的值即该单词首字母在order中的位置，res的索引为该单词在words中的位置
func isAlienSorted(words []string, order string) bool {
	res := make([]int, 0, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(order); j++ {
			if words[i][0] == order[j] {
				res = append(res, j)
				break
			}
		}
	}
	tmp := math.MinInt64
	for i := 0; i < len(res); i++ {
		if tmp < res[i] {
			tmp = res[i]
		} else if tmp == res[i] {
			//如果两个res相邻的两个值相同，说明这两个单词首字母相同，此时调用check逐位检查
			if check(words[i-1], words[i], order) {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func check(str1, str2, order string) bool {
	str1Position := math.MinInt8
	str2Position := math.MinInt8
	str1Size := len(str1)
	str2Size := len(str2)
	loop := min(str1Size, str2Size)
	for i := 1; i < loop; i++ {
		if str1[i] == str2[i] {
			continue
		} else {
			for k := 0; k < len(order); k++ {
				if str1[i] == order[k] {
					str1Position = k
				}
				if str2[i] == order[k] {
					str2Position = k
				}
				if str1Position < str2Position && str1Position != math.MinInt8 && str2Position != math.MinInt8 {
					return true
				} else if str1Position > str2Position && str1Position != math.MinInt8 && str2Position != math.MinInt8 {
					return false
				}
			}
		}
	}
	//如果str1和str2中较短的那个是另一个的真子集，且str1的长度小于str2则返回true
	if len(str1) <= len(str2) {
		return true
	} else {
		return false
	}
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
