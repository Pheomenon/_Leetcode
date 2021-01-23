package main

import "fmt"

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：

num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库，也不能直接将输入的字符串转换为整数形式
通过次数91,109提交次数175,654

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	num1 := "9"
	num2 := "99"
	fmt.Println(addStrings(num1, num2))
}

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	var ret byte
	ret = '0'
	var ans []byte

	if i > j {
		ans = make([]byte, len(num1))
	} else {
		ans = make([]byte, len(num2))
	}

	for i > -1 && j > -1 {
		if len(num1) >= len(num2) {
			ans[i], ret = add(num1[i]-'0', num2[j]-'0'+ret-'0')
			i--
			j--
		} else {
			ans[j], ret = add(num1[i]-'0', num2[j]-'0'+ret-'0')
			i--
			j--
		}
	}

	if i == j {
		if ret != '0' {
			ans = append([]byte{ret}, ans...)
		}
		return string(ans)
	}

	if i != -1 {
		for ; i >= 0; i-- {
			ans[i], ret = add(num1[i]-'0', ret-'0')
		}
	}

	if j != -1 {
		for ; j >= 0; j-- {
			ans[j], ret = add(num2[j]-'0', ret-'0')
		}
	}
	if ret != '0' {
		ans = append([]byte{ret}, ans...)
	}

	return string(ans)
}

func add(x, y byte) (byte, byte) {
	sum := x + y
	if sum >= 10 {
		return sum - 10 + '0', sum/10 + '0'
	} else {
		return sum + '0', 0 + '0'
	}
}
