package leetcode

/*
个人思路：一个map存储key，一个list存储顺序，但是go的list包用的很少，不熟悉
答案：可以自己实现一个双向链表，也不难。node存储prev和next，记录时存储head和tail
*/

type LRUCache struct {
	cache      map[int]*Node
	head, tail *Node
	cap        int
}

type Node struct {
	Key, Value int
	next, prev *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cache: make(map[int]*Node),
		cap:   capacity,
	}
}

func (this *LRUCache) Add(node *Node) {
	node.next = this.head
	if this.head != nil {
		this.head.prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) Remove(node *Node) {
	if node == this.head {
		if node.next != nil {
			node.next.prev = nil
		}
		this.head = node.next
		node.next = nil
		return
	}
	if node == this.tail {
		node.prev.next = nil
		this.tail = node.prev
		node.prev = nil
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.Remove(node)
		this.Add(node)
		return node.Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		this.Remove(node)
		node.Value = value
		this.Add(node)
	} else {
		node = &Node{Key: key, Value: value}
		this.Add(node)
		this.cache[key] = node
		if len(this.cache) > this.cap {
			delete(this.cache, this.tail.Key)
			this.Remove(this.tail)
		}
	}
}
