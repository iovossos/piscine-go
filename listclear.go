// Write a function ListClear that deletes all nodes from a linked list l.

// Tip: assign the list's pointer to nil.
package piscine

func ListClear(l *List) {
	for l.Head != nil {
		temp := l.Head
		l.Head = l.Head.Next
		temp.Next = nil
	}

	l.Head = nil
	l.Tail = nil
}
