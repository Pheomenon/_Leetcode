/**
运用你所掌握的数据结构，设计和实现一个 LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字已经存在，则变更其数据值；
如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，
它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。

进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例：
输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2],       [1, 1], [2, 2], [1],  [3, 3], [2],  [4, 4], [1],   [3],   [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

[null, null, null, 1, null, 2, null, -1, 3, 4]
解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

提示：
1 <= capacity <= 3000
0 <= key <= 3000
0 <= value <= 104
最多调用 3 * 104 次 get 和 put
通过次数110,952提交次数217,144

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

func main() {
	op := Constructor(2)
	op.Put(2, 1)
	op.Put(2, 2)
	op.Get(2)
	op.Put(1, 1)
	op.Put(4, 1)
	op.Get(2)
}

type LRUCache struct {
	capacity int
	cache    map[int]*node
	head     *node
	tail     *node
	len      int
}

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}

func (c *LRUCache) addFirst(item *node) {
	item.next = c.head.next
	item.prev = c.head
	c.head.next.prev = item
	c.head.next = item
}

func (c *LRUCache) remove(item *node) {
	item.next.prev = item.prev
	item.prev.next = item.next
}

func (c *LRUCache) removeLast() *node {
	node := c.tail.prev
	c.remove(node)
	return node
}

func (c *LRUCache) size() int {
	return c.len
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*node{},
		head:     &node{},
		tail:     &node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.cache[key]; !ok {
		return -1
	} else {
		c.Put(key, v.value)
		return v.value
	}
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; !ok {
		node := &node{
			key:   key,
			value: value,
			next:  nil,
			prev:  nil,
		}
		c.addFirst(node)
		c.cache[key] = node
		c.len++
		if c.capacity < c.len {
			ret := c.removeLast()
			delete(c.cache, ret.key)
			c.len--
		}
	} else {
		newNode := c.cache[key]
		newNode.value = value
		c.remove(newNode)
		c.addFirst(newNode)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
