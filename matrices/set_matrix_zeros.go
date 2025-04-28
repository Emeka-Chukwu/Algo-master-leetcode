package main

import "fmt"

// Description
// Editorial
// Editorial
// Solutions
// Solutions
// Submissions
// Submissions

// Code
// Testcase
// Test Result
// Test Result
// 73. Set Matrix Zeroes
// Medium
// Topics
// Companies
// Hint
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

// You must do it in place.

// Example 1:

// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
// Example 2:

// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Constraints:

// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -231 <= matrix[i][j] <= 231 - 1

// Follow up:

// A straightforward solution using O(mn) space is probably a bad idea.
// A simple improvement uses O(m + n) space, but still not the best solution.
// Could you devise a constant space solution?
// Accepted
// 1.9M
// Submissions
// 3.2M
// Acceptance Rate
// 59.4%
// Topics
// Companies
// Hint 1
// Hint 2
// Hint 3
// Hint 4
// Similar Questions
// Discussion (149)

// solutin is time complexity is 2(m*n )

func setZeroes(matrix [][]int) [][]int {
	results := make([][]int, len(matrix))
	for i := range matrix {
		results[i] = make([]int, len(matrix[0]))
	}
	var existingRow map[int]bool = make(map[int]bool)
	var existingCol map[int]bool = make(map[int]bool)

	for indexCol, rows := range matrix {
		for indexRow, row := range rows {
			if row == 0 {
				existingCol[indexCol] = true
				existingRow[indexRow] = true
			}
		}
	}

	for indexCol, rows := range matrix {
		if existingCol[indexCol] {
			continue
		}
		for indexRow, row := range rows {
			if existingRow[indexRow] {
				continue
			}
			results[indexCol][indexRow] = row

		}
	}
	return results
}

func main() {
	fmt.Println(setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}))
	fmt.Println(setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}))
}

// 1
// // Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// // Output: [[1,0,1],[0,0,0],[1,0,1]]
// 2
// // Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// // Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
