package main

import "fmt"

/*
请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
实现词典类 WordDictionary ：
WordDictionary() 初始化词典对象
void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true；
否则，返回 false 。word 中可能包含一些 '.' ，每个. 都可以表示任何一个字母。

示例：
输入：
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
输出：
[null,null,null,null,false,true,true,true]

解释：
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True

提示：

1 <= word.length <= 500
addWord 中的 word 由小写英文字母组成
search 中的 word 由 '.' 或小写英文字母组成
最调用多 50000 次 addWord 和 search

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-add-and-search-words-data-structure
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	wordDictionary.Search("pad")              // return False
	wordDictionary.Search("bad")              // return True
	fmt.Println(wordDictionary.Search(".ad")) // return True
	wordDictionary.Search("b..")              // return True

}

type WordDictionary struct {
	isEnd bool
	next  [26]*WordDictionary
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = &WordDictionary{}
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	var searchHelp func(word string, node *WordDictionary) bool
	searchHelp = func(word string, node *WordDictionary) bool {
		for i := 0; i < len(word); i++ {
			if word[i] == '.' {
				for _, dictionary := range node.next {
					if dictionary != nil {
						if searchHelp(word[i+1:], dictionary) {
							return true
						}
					}
				}
				return false
			}
			if node.next[word[i]-'a'] == nil {
				return false
			} else {
				node = node.next[word[i]-'a']
			}
		}
		return node.isEnd == true
	}
	return searchHelp(word, this)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
