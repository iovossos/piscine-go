package piscine

func ListSize(l *List) int {
	size := 0
	for node := l.Head; node != nil; node = node.Next {
		size++
	}
	return size
}
