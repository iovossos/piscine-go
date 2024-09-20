package main

import "github.com/01-edu/z01"

// Define constants for door states
const (
	OPEN  = true
	CLOSE = false
)

// Define a Door struct
type Door struct {
	state bool
}

// PrintStr prints a string rune by rune
func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// CloseDoor sets the door's state to closed
func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = CLOSE
	return true
}

// OpenDoor sets the door's state to open
func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = OPEN
	return true
}

// IsDoorOpen checks if the door is open
func IsDoorOpen(door *Door) bool {
	PrintStr("Is the Door opened?")
	return door.state == OPEN
}

// IsDoorClose checks if the door is closed
func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Is the Door closed?")
	return ptrDoor.state == CLOSE
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
