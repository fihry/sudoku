// created by melfihry pooler in z01
// date: 2024/01/21
// description: a program that solves a sudoku puzzle
// usage: go run sudoku.go "test" "test" "test" "test" "test" "test" "test" "test" "test" | cat -e
// test1 : "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79" | cat -e
// test2 : 8........ ..36..... .7..9.2.. .5...7... ....457.. ...1...3. ..1....68 ..85...1. .9....4.. | cat -e
// test3 : . 1 2 3 4 | cat -e            // Error
// created by melfihry pooler in z01oujda
// date: 2024/01/21
// description: a program that solves a sudoku puzzle
// usage: go run sudoku.go "test" "test" "test" "test" "test" "test" "test" "test" "test" | cat -e
// test1 : "53..7...." "6..195..." ".98....6." "8...6...3" "4..8.3..1" "7...2...6" ".6....28." "...419..5" "....8..79" | cat -e
// test2 : 8........ ..36..... .7..9.2.. .5...7... ....457.. ...1...3. ..1....68 ..85...1. .9....4.. | cat -e
// test3 : . 1 2 3 4 | cat -e            // Error
package main

import (
	"os"

	"github.com/01-edu/z01"
)

const N = 9

// Check if the number is safe
func isSafe(board [][]int, row int, col int, num int) bool {
	// Check the row
	for i := 0; i < N; i++ {
		if board[row][i] == num {
			return false
		}
	}
	// Check the column
	for i := 0; i < N; i++ {
		if board[i][col] == num {
			return false
		}
	}
	// Check the 3x3 grid
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i+startRow][j+startCol] == num {
				return false
			}
		}
	}
	return true
}

// Solve the sudoku
func solveSudoku(board [][]int, row int, col int) bool {
	if col == N {
		col = 0
		row++
		if row == N {
			return true
		}
	}
	if board[row][col] == 0 {
		for num := 1; num <= N; num++ {
			if isSafe(board, row, col, num) {
				board[row][col] = num
				if solveSudoku(board, row, col+1) {
					return true
				}
				board[row][col] = 0
			}
		}
	} else {
		return solveSudoku(board, row, col+1)
	}
	return false
}

// Check if the sudoku is valid
func isValidSudoku(board [][]int) bool {
	for i := 0; i < N; i++ {
		row := make([]bool, N+1)
		col := make([]bool, N+1)
		for j := 0; j < N; j++ {
			if board[i][j] != 0 {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}
			if board[j][i] != 0 {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}
		}
	}
	for i := 0; i < N; i += 3 {
		for j := 0; j < N; j += 3 {
			block := make([]bool, N+1)
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					if board[i+k][j+l] != 0 {
						if block[board[i+k][j+l]] {
							return false
						}
						block[board[i+k][j+l]] = true
					}
				}
			}
		}
	}
	return true
}

// Print the board
func printBoard(board [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if board[i][j] != 0 {
				z01.PrintRune(rune(board[i][j] + '0'))
			} else {
				z01.PrintRune('.')
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

// Print Error
func printError() {
	for _, i := range "Error" {
		z01.PrintRune(rune(i))
	}
	z01.PrintRune('\n')
}

func main() {
	board := make([][]int, N)
	for i := range board {
		board[i] = make([]int, N)
	}
	if len(os.Args) != N+1 {
		printError()
		return
	}
	for i := 1; i <= N; i++ {
		if len(os.Args[i]) != N {
			printError()
			return
		}
		for j := 0; j < N; j++ {
			if os.Args[i][j] == '.' {
				board[i-1][j] = 0
			} else if os.Args[i][j] >= '1' && os.Args[i][j] <= '9' {
				board[i-1][j] = int(os.Args[i][j] - '0')
			} else {
				printError()
				return
			}
		}
	}
	if !isValidSudoku(board) {
		printError()
		return
	}
	if !solveSudoku(board, 0, 0) {
		printError()
		return
	}
	printBoard(board)
}
