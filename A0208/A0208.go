package main

/*
实现一个 Trie (前缀树)，包含insert,search, 和startsWith这三个操作。

示例:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // 返回 true
trie.search("app");     // 返回 false
trie.startsWith("app"); // 返回 true
trie.insert("app");
trie.search("app");     // 返回 true
说明:

你可以假设所有的输入都是由小写字母a-z构成的。
保证所有输入均为非空字符串。
通过次数66,933提交次数96,320

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-trie-prefix-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")   // 返回 true
	trie.Search("app")     // 返回 false
	trie.StartsWith("app") // 返回 true
	trie.Insert("app")
	trie.Search("app") // 返回 true
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, w := range word {
		if this.next[w-'a'] == nil {
			this.next[w-'a'] = &Trie{}
		}
		this = this.next[w-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, w := range word {
		if this.next[w-'a'] == nil {
			return false
		} else {
			this = this.next[w-'a']
		}
	}
	return this.isEnd == true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, p := range prefix {
		if this.next[p-'a'] == nil {
			return false
		} else {
			this = this.next[p-'a']
		}
	}
	return true
}

/**
* Your Trie object will be instantiated and called as such:
* obj := Constructor();
* obj.Insert(word);
* param_2 := obj.Search(word);
* param_3 := obj.StartsWith(prefix);
 */
