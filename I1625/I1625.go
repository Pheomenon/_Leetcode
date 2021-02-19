package main

/*
设计和构建一个“最近最少使用”缓存，该缓存会删除最近最少使用的项目。缓存应该从键映射到值(允许你插入和检索特定键对应的值)，并在初始化时指定最大容量。当缓存被填满时，它应该删除最近最少使用的项目。

它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

示例:

LRUCache cache = new LRUCache(2); // 缓存容量

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // 返回  1
cache.put(3, 3);    // 该操作会使得密钥 2 作废
cache.get(2);       // 返回 -1 (未找到)
cache.put(4, 4);    // 该操作会使得密钥 1 作废
cache.get(1);       // 返回 -1 (未找到)
cache.get(3);       // 返回  3
cache.get(4);       // 返回  4
通过次数14,796提交次数27,578

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lru-cache-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	cache := Constructor(2) // 缓存容量

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Get(1)    // 返回  1
	cache.Put(3, 3) // 该操作会使得密钥 2 作废
	cache.Get(2)    // 返回 -1 (未找到)
	cache.Put(4, 4) // 该操作会使得密钥 1 作废
	cache.Get(1)    // 返回 -1 (未找到)
	cache.Get(3)    // 返回  3
	cache.Get(4)
}

type node struct {
	key, val   int
	next, prev *node
}

type LRUCache struct {
	cache    map[int]*node
	capacity int
	len      int
	head     *node
	tail     *node
}

func (l *LRUCache) addToFirst(n *node) {
	n.next = l.head.next
	n.prev = l.head
	l.head.next.prev = n
	l.head.next = n
}

func (l *LRUCache) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *LRUCache) removeLast() *node {
	node := l.tail.prev
	l.remove(node)
	return node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*node{},
		capacity: capacity,
		head:     &node{},
		tail:     &node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	} else {
		n := this.cache[key]
		this.remove(n)
		this.addToFirst(n)
		value := n.val
		return value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		newNode := &node{
			key:  key,
			val:  value,
			next: &node{},
			prev: &node{},
		}
		this.addToFirst(newNode)
		this.cache[key] = newNode
		this.len++
		if this.len > this.capacity {
			n := this.removeLast()
			delete(this.cache, n.key)
			this.len--
		}
	} else {
		this.cache[key].val = value
		this.remove(this.cache[key])
		this.addToFirst(this.cache[key])

	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
