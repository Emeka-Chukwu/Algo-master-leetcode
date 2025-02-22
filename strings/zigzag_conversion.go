package main

import "fmt"

// 6. Zigzag Conversion
// Medium
// Topics
// Companies
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

// Constraints:

func convert(s string, numRows int) string {
	resultMap := map[int]string{}
	result := ""
	count := 0
	addition := true
	for i := 0; i < len(s); i++ {
		if count == 0 {
			addition = true
		} else if count == numRows-1 {
			addition = false
		}
		resultMap[count] += string(s[i])
		if addition {
			count++
		} else {
			count--
		}
	}
	for i := 0; i < numRows; i++ {
		result += resultMap[i]
	}
	return result
}

func main() {

	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))

}
