// Write a function ListSize that returns the number of elements in a linked list l.
package piscine

func ListSize(l *List) int {
	size := 0
	for node := l.Head; node != nil; node = node.Next {
		size++
	}
	return size
}
