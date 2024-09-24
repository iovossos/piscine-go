package piscine

type NodeL1 struct {
	Data interface{}
	Next *NodeL
}

type List1 struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List1, data interface{}) {
	newNode := &NodeL{Data: data}

	newNode.Next = l.Head

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}
}
