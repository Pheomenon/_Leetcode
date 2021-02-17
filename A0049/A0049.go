package main

/*

给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
通过次数164,198提交次数250,821
*/

func main() {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	groupAnagrams(strs)
}

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	m := map[[26]uint8][]string{}
	ans := make([][]string, 0)

	for i := 0; i < len(strs); i++ {
		cnt := [26]uint8{}
		for _, v := range strs[i] {
			cnt[v-'a']++
		}
		m[cnt] = append(m[cnt], strs[i])
	}

	for _, strings := range m {
		ans = append(ans, strings)
	}
	return ans
}
