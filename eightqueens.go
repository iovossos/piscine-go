package piscine

import "fmt"

// Helper function to check if the current queen placement is valid
func isValid(board []int, row, col int) bool {
	for prevRow := 0; prevRow < row; prevRow++ {
		// Check column and diagonals
		if board[prevRow] == col || board[prevRow]-prevRow == col-row || board[prevRow]+prevRow == col+row {
			return false
		}
	}
	return true
}

// Recursive function to solve the Eight Queens puzzle
func solveQueens(board []int, row int) {
	if row == len(board) {
		// Print the solution
		for _, col := range board {
			fmt.Print(col + 1)
		}
		fmt.Println()
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row] = col
			solveQueens(board, row+1)
		}
	}
}

// Function to start the Eight Queens puzzle solver
func EightQueens() {
	board := make([]int, 8) // Create an array to represent the board
	solveQueens(board, 0)   // Start solving from the first row
}
