/**
在一个由 'L' , 'R' 和 'X' 三个字符组成的字符串（例如"RXXLRXRXL"）中进行移动操作。
一次移动操作指用一个"LX"替换一个"XL"，或者用一个"XR"替换一个"RX"。
现给定起始字符串start和结束字符串end，请编写代码，
当且仅当存在一系列移动操作使得start可以转换成end时， 返回True。

示例 :
输入: start = "RXXLRXRXL", end = "XRLXXRRLX"
输出: True
解释:
我们可以通过以下几步将start转换成end:
RXXLRXRXL ->
XRXLRXRXL ->
XRLXRXRXL ->
XRLXXRRXL ->
XRLXXRRLX

提示：
1 <= len(start) = len(end) <= 10000。
start和end中的字符串仅限于'L', 'R'和'X'。
通过次数4,018提交次数12,598

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-adjacent-in-lr-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	start := "XLLR"
	end := "LXLX"
	fmt.Println(canTransform(start, end))
}

func canTransform(start string, end string) bool {
	if len(start) == 1 {
		return start == end
	}
	S := make([]string, 0, len(start))
	E := make([]string, 0, len(end))
	for _, v := range start {
		S = append(S, string(v))
	}
	for _, v := range end {
		E = append(E, string(v))
	}
	for i := 0; i < len(S)-1; i++ {
		if S[i] != E[i] {
			if S[i+1] == "R" && S[i] == "X" || S[i+1] == "X" && S[i] == "R" || S[i+1] == "X" && S[i] == "L" || S[i+1] == "L" && S[i] == "X" {
				S[i+1], S[i] = S[i], S[i+1]
			}
			if E[i+1] == S[i+1] && E[i] == S[i] {
				continue
			} else {
				return false
			}
		} else {
			continue
		}
	}
}
