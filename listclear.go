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
