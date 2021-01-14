package main

/*
给定由若干0和1组成的数组 A。我们定义N_i：从A[0] 到A[i]的第 i个子数组被解释为一个二进制数（从最高有效位到最低有效位）。
返回布尔值列表answer，只有当N_i可以被 5整除时，答案answer[i] 为true，否则为 false。

示例 1：
输入：[0,1,1]
输出：[true,false,false]
解释：
输入数字为 0, 01, 011；也就是十进制中的 0, 1, 3 。只有第一个数可以被 5 整除，因此 answer[0] 为真。

示例 2：
输入：[1,1,1]
输出：[false,false,false]

示例 3：
输入：[0,1,1,1,1,1]
输出：[true,false,false,false,true,false]

示例 4：
输入：[1,1,1,0,1]
输出：[false,false,false,false,false]

提示：

1 <= A.length <= 30000
A[i] 为 0 或 1
通过次数19,835提交次数39,338

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-prefix-divisible-by-5
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	num := []int{0, 1, 1}
	prefixesDivBy5(num)
}

func prefixesDivBy5(A []int) []bool {
	ret := 0
	ans := make([]bool, len(A))
	for i, v := range A {
		ret = (ret<<1 | v) % 5
		if ret == 0 {
			ans[i] = true
			continue
		}
	}
	return ans
}
