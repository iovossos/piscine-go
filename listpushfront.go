package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL1
}

type List1 struct {
	Head *NodeL1
	Tail *NodeL1
}

func ListPushFront(l *List1, data interface{}) {
	newNode := &NodeL1{Data: data}

	newNode.Next = l.Head

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}
}
