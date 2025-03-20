package main

// 49. Group Anagrams
// Medium
// Topics
// Companies
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]

// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Explanation:

// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.
// Example 2:

// Input: strs = [""]

// Output: [[""]]

// Example 3:

// Input: strs = ["a"]

// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.

func groupAnagrams(strs []string) [][]string {

	var results [][]string = make([][]string, 0)
	var resultMap map[int][]string = map[int][]string{}

	for _, group := range strs {
		numTotal := 0
		for i := 0; i < len(group); i++ {
			numTotal = numTotal + int(group[i])
		}
		resultMap[numTotal] = append(resultMap[numTotal], group)
	}
	for _, value := range resultMap {
		results = append(results, value)
	}
	return results
}

// func main() {
// 	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
// 	fmt.Println(groupAnagrams([]string{""}))
// 	fmt.Println(groupAnagrams([]string{"a"}))
// }
