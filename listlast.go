// Write a function ListLast that returns the Data interface of the last element of a linked list l.
package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	node := l.Head

	for node.Next != nil {
		node = node.Next
	}

	return node.Data
}
