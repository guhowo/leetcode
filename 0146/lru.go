package _146

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Capacity int
	Index    map[int]*Node
	Head     *Node
	Tail     *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		Capacity: capacity,
		Index:    make(map[int]*Node),
		Head:     &Node{},
		Tail:     &Node{},
	}
	cache.Head.Next = cache.Tail
	cache.Tail.Prev = cache.Head
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Index[key]
	if !ok {
		return -1
	}
	this.Remove(node)
	this.SetHead(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Index[key]
	if ok {
		this.Remove(node)
		node.Val = value
	} else {
		if this.Capacity == len(this.Index) {
			delete(this.Index, key)
			this.Remove(this.Tail.Prev)
		}

	}
	this.SetHead(node)
}

func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) SetHead(node *Node) {
	node.Next = this.Head.Next
	node.Next.Prev = node
	this.Head.Next = node
	node.Prev = this.Head
}
