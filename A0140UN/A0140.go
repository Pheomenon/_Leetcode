package main

/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，
在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

说明：
分隔时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

示例 1：
输入:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
输出:
[
 "cats and dog",
 "cat sand dog"
]

示例 2：
输入:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
输出:
[
 "pine apple pen apple",
 "pineapple pen apple",
 "pine applepen apple"
]
解释: 注意你可以重复使用字典中的单词。
示例 3：

输入:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
输出:
[]
通过次数39,615提交次数89,382

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-break-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {

}

//type trie struct {
//	isEnd bool
//	next  [26]*trie
//}
//
//func (t *trie) insert(word string) {
//	for _, v := range word {
//		if t.next[v-'a'] == nil {
//			t.next[v-'a'] = &trie{}
//		}
//		t = t.next[v-'a']
//	}
//	t.isEnd = true
//}
//
//func (t *trie) search(word string) []string {
//	texts := make([]string, 0)
//	for i, v := range word {
//		if t.next[v-'a'] == nil {
//			break
//		} else {
//			if t.isEnd {
//				texts = append(texts, word[:i+1])
//			}
//			t = t.next[v-'a']
//		}
//	}
//	return texts
//}
//
//func wordBreak(s string, wordDict []string) []string {
//	trie := &trie{}
//	texts := make([][]string, 0)
//	for _, v := range wordDict {
//		trie.insert(v)
//	}
//	testI := 0
//	for len(s) > 0 {
//		searchRes := trie.search(s)
//		for i := 0; i < len(searchRes); i++ {
//			texts[testI] = make([]string, len(searchRes))
//			texts[testI] = append(texts[testI], searchRes...)
//		}
//		testI++
//		min := math.MaxInt64
//		for _, v := range searchRes {
//			if len(v) < min {
//				min = len(v)
//			}
//		}
//		s = s[min:]
//	}
//	ans := make([]string, 0)
//
//}
//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//
//	return y
//}
