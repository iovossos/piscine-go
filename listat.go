// Write a function ListAt that takes a pointer to the head of the list l and an int pos as parameters. This function should return the pointer to the NodeL in the position pos of the linked list l.

// In case of error the function should return nil.
package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; l != nil; i++ {
		if i == pos {
			return l
		}
		l = l.Next
	}
	return nil
}
