// Write a function ListReverse that reverses the order of the elements of a given linked list l.
package piscine

func ListReverse(l *List) {
	var prev *NodeL
	current := l.Head
	var next *NodeL

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	l.Tail = l.Head
	l.Head = prev
}
