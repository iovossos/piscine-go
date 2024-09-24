// Write a function ListPushFront that inserts a new element NodeL at the beginning of the list l while using the structure List
package piscine

func ListPushFront(l *List, data interface{}) {
	newNode := &NodeL{Data: data}

	newNode.Next = l.Head

	l.Head = newNode

	if l.Tail == nil {
		l.Tail = newNode
	}
}
