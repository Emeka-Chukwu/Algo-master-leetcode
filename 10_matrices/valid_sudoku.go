package main

import "fmt"

// Code
// Testcase
// Test Result
// Test Result
// 36. Valid Sudoku
// Medium
// Topics
// Companies
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Example 1:

// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
// Example 2:

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6","3",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

// Constraints:

// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.
// Accepted
// 2M
// Submissions
// 3.1M
// Acceptance Rate
// 62.0%
// Topics
// Companies
// Similar Questions
// Discussion (194)

func isValidSudoku(board [][]byte) bool {

	var isDuplicate map[byte]bool = make(map[byte]bool)
	for _, col := range board {
		isDuplicate = make(map[byte]bool)
		for _, row := range col {
			if isDuplicate[row] {
				return false
			}
			if row == 46 {
				continue
			}
			isDuplicate[row] = true
		}
	}
	for indexCol, col := range board {
		isDuplicate = make(map[byte]bool)
		for rowIndex, _ := range col {
			if isDuplicate[board[rowIndex][indexCol]] {
				return false
			}

			if board[rowIndex][indexCol] == 46 {
				continue
			}
			isDuplicate[board[rowIndex][indexCol]] = true
		}
	}

	startCol := 0
	startRow := 3
	colIncrement := 0
	rowIncrement := 0
	isDuplicate = make(map[byte]bool)

	for {
		if isDuplicate[board[colIncrement][rowIncrement]] {
			return false
		}
		isDuplicate[board[colIncrement][rowIncrement]] = true
		fmt.Println(colIncrement)

		if rowIncrement%9 == 0 && colIncrement%9 == 0 {
			return true
		} else if colIncrement%9 == 0 {
			colIncrement = startCol
		} else if colIncrement%3 == 0 && rowIncrement%3 == 0 {
			isDuplicate = make(map[byte]bool)

			rowIncrement = startRow
			break
		}
		if rowIncrement%3 == 0 {
			rowIncrement = startRow
			colIncrement++
			break
		}
		rowIncrement++

	}

	return true

}

func main() {
	board1 := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	board2 := [][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '3', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board1))
	fmt.Println(isValidSudoku(board2))
	fmt.Println()

	fmt.Println(isValidSudoku1(board1))
	fmt.Println(isValidSudoku1(board2))
}

func isValidSudoku1(board [][]byte) bool {

	var rows []map[byte]bool = make([]map[byte]bool, 9)
	var cols []map[byte]bool = make([]map[byte]bool, 9)
	var boxIndex []map[byte]bool = make([]map[byte]bool, 9)

	for i := 0; i < len(board); i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		boxIndex[i] = map[byte]bool{}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}
			boxPos := (i/3)*3 + j/3
			if cols[i][num] || rows[j][num] || boxIndex[boxPos][num] {
				// fmt.Println(cols[i][num], rows[j][num], boxIndex[boxPos][num])
				return false
			}

			cols[i][num] = true
			rows[j][num] = true
			boxIndex[boxPos][num] = true
		}
	}
	return true

}
