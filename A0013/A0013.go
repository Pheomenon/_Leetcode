package main

import "fmt"

/*
罗马数字包含以下七种字符: I，V，X，L，C，D和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做XII，即为X+II。 27 写做XXVII, 即为XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C(100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示400 和 900。
给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

示例 1:

输入: "III"
输出: 3
示例 2:

输入: "IV"
输出: 4
示例 3:

输入: "IX"
输出: 9
示例 4:

输入: "LVIII"
输出: 58
解释: L = 50, V= 5, III = 3.
示例 5:

输入: "MCMXCIV"
输出: 1994
解释: M = 1000, CM = 900, XC = 90, IV = 4.

提示：

题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
IC 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
关于罗马数字的详尽书写规则，可以参考 罗马数字 - Mathematics 。
通过次数302,482提交次数485,463

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/roman-to-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

func romanToInt(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		tmp := s[i]
		if i < len(s)-1 {
			if s[i+1] == 'V' && s[i] == 'I' || s[i+1] == 'X' && s[i] == 'I' {
				ans -= I
				continue
			}
			if s[i+1] == 'L' && s[i] == 'X' || s[i+1] == 'C' && s[i] == 'X' {
				ans -= X
				continue
			}
			if s[i+1] == 'D' && s[i] == 'C' || s[i+1] == 'M' && s[i] == 'C' {
				ans -= C
				continue
			}
		}
		if tmp == 'M' {
			ans += M
		} else if tmp == 'D' {
			ans += D
		} else if tmp == 'C' {
			ans += C
		} else if tmp == 'L' {
			ans += L
		} else if tmp == 'X' {
			ans += X
		} else if tmp == 'V' {
			ans += V
		} else if tmp == 'I' {
			ans += I
		}
	}
	return ans
}
