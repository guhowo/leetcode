package _146

type Node struct {
	Element int
	Next    *Node
	Prev    *Node
}

type List struct {
	Head   *Node //哨兵，不存数据
	Tail   *Node
	Length int
}

type LRUCache struct {
	Capacity int
	Index    map[int]*Node
	List     *List
}

func Constructor(capacity int) LRUCache {
	head := &Node{Next: nil, Prev: nil}
	return LRUCache{
		Capacity: capacity,
		Index:    make(map[int]*Node),
		List:     &List{Head: head, Tail: head, Length: 0},
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Index[key]; !ok {
		return -1
	} else {
		head := this.List.Head
		p := this.Index[key]
		p.Prev.Next = p.Next
		if p.Next != nil {
			p.Next.Prev = p.Prev
		} else {
			this.List.Tail = p.Prev
		}
		p.Next = head.Next
		head.Next.Prev = p
		head.Next = p
		p.Prev = head
		return this.Index[key].Element
	}
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{
		Element: value,
		Next:    nil,
		Prev:    this.List.Tail,
	}
	head := this.List.Head
	if _, ok := this.Index[key]; ok {
		p := this.Index[key]
		p.Prev.Next = p.Next
		if p.Next != nil {
			p.Next.Prev = p.Prev
		} else {
			this.List.Tail = p.Prev
		}
		p.Next = head.Next
		head.Next.Prev = p
		head.Next = p
		p.Prev = head
	} else if len(this.Index)+1 <= this.Capacity {
		if head.Next == nil {
			this.List.Tail = node
			head.Next = node
			node.Prev = head
			return
		}
		node.Next = head.Next
		head.Next.Prev = node
		node.Prev = head
		head.Next = node
		this.Index[key] = node
	} else {
		node.Next = head.Next
		head.Next.Prev = node
		node.Prev = head
		head.Next = node
		delete(this.Index, this.List.Tail.Element)
		this.List.Tail.Prev.Next = nil
		this.List.Tail = this.List.Tail.Prev
	}
}
