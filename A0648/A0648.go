package main

import (
	"fmt"
	"math"
	"strings"
)

/*
在英语中，我们有一个叫做词根(root)的概念，它可以跟着其他一些词组成另一个较长的单词——我们称这个词为继承词(successor)。
例如，词根an，跟随着单词other(其他)，可以形成新的单词another(另一个)。
现在，给定一个由许多词根组成的词典和一个句子。你需要将句子中的所有继承词用词根替换掉。
如果继承词有许多可以形成它的词根，则用最短的词根替换它。
你需要输出替换之后的句子。

示例 1：
输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

示例 2：
输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
输出："a a b c"

示例 3：
输入：dictionary = ["a", "aa", "aaa", "aaaa"], sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
输出："a a a a a a a a bbb baba a"

示例 4：
输入：dictionary = ["catt","cat","bat","rat"], sentence = "the cattle was rattled by the battery"
输出："the cat was rat by the bat"

示例 5：
输入：dictionary = ["ac","ab"], sentence = "it is abnormal that this solution is accepted"
输出："it is ab that this solution is ac"

提示：
1 <= dictionary.length<= 1000
1 <= dictionary[i].length <= 100
dictionary[i]仅由小写字母组成。
1 <= sentence.length <= 10^6
sentence仅由小写字母和空格组成。
sentence 中单词的总量在范围 [1, 1000] 内。
sentence 中每个单词的长度在范围 [1, 1000] 内。
sentence 中单词之间由一个空格隔开。
sentence 没有前导或尾随空格。
通过次数13,938提交次数24,905

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/replace-words
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	dict := []string{
		"cat", "bat", "rat",
	}
	sentence := "the cattle was rattled by the battery"
	fmt.Println(replaceWords(dict, sentence))
}

type trie struct {
	isEnd bool
	next  [26]*trie
}

func (t *trie) insert(word string) {
	for _, v := range word {
		if t.next[v-'a'] == nil { //如果该字母不在字典树中，则构建一个新节点
			t.next[v-'a'] = &trie{}
		}
		t = t.next[v-'a'] //如果该字母在字典树中存在则指针进入下一层
	}
	t.isEnd = true //标记最后一个节点
}

func (t *trie) search(word string) string {
	ans := make([]string, 0)
	minL := math.MaxInt64
	for i := 0; i < 1; i++ {
		//这里可以用循环也可以不用，如果i的上限为len(word)那就是从整个单词里找有没有字典中存在的词，
		//根据题意只能找word前缀匹配的情况，所以i的上限为1
		curr := t //这里保存根节点指针，如果只匹配前缀，这一步可以忽略
		res := make([]byte, 0)
		if curr.next[word[i]-'a'] == nil {
			continue
		}
		for j := i; j < len(word); j++ {
			if curr != nil {
				if curr.isEnd {
					ans = append(ans, string(res))
					break
				}
				res = append(res, (word[j]))
				curr = curr.next[word[j]-'a']
			}
		}
	}
	for i, an := range ans { //找出符合的最短前缀
		if len(an) < minL {
			minL = i
		}
	}
	if minL != math.MaxInt64 {
		return ans[minL]
	}
	return ""
}

func replaceWords(dictionary []string, sentence string) string {
	trie := &trie{}
	for _, s := range dictionary {
		trie.insert(s)
	}

	context := strings.Split(sentence, " ")
	for i := 0; i < len(context); i++ {
		searchRes := trie.search(context[i])
		if searchRes != "" {
			context[i] = searchRes
		}
	}
	return strings.Join(context, " ")
}
