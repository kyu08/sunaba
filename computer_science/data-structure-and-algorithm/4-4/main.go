package main

func main() {
	firstNode := NewNode()
	firstNode.insert(1)
}

type Node struct {
	key  int
	prev *Node
	next *Node
}

// NewNode 番兵(データの起点を表すデータ)を返す
// prev, nextに自身をさすポインタを持つ
func NewNode() *Node {
	nilNode := &Node{}
	nilNode.prev = nilNode
	nilNode.next = nilNode
	return nilNode
}

// insert 番兵の次にデータを挿入する
func (n *Node) insert(val int) {
	newNode := &Node{key: val}
	n.next.prev = newNode
	newNode.next = n.next
	n.next = newNode
	newNode.prev = n
}

// listSearch
func (n *Node) listSearch(key int) *Node {
	now := n.next
	for {
		if now == n {
			return nil
		}
		if now.key == key {
			return now
		}
		now = now.next
	}
}
