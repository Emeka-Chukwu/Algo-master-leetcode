package main

// Code
// Testcase
// Test Result
// Test Result
// 122. Best Time to Buy and Sell Stock II
// Medium
// Topics
// Companies
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.

// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day.

// Find and return the maximum profit you can achieve.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 7
// Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
// Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
// Total profit is 4 + 3 = 7.
// Example 2:

// Input: prices = [1,2,3,4,5]
// Output: 4
// Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
// Total profit is 4.
// Example 3:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.

// Constraints:

// 1 <= prices.length <= 3 * 104
// 0 <= prices[i] <= 104
// / easy stock I
func maxProfit(prices []int) int {
	profit, left, right := 0, 0, 1

	for i := 1; i < len(prices); i++ {
		if prices[left] > prices[right] {
			left++
		} else {
			diff := prices[right] - prices[left]
			if diff > profit {
				profit = diff
			}
		}
		right++
	}
	return profit
}

// ////// STOCK II
func maxProfit1(prices []int) int {
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// func main() {
// 	fmt.Println(" Best Time to Buy and Sell Stock I")
// 	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
// 	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
// 	fmt.Println()
// 	fmt.Println(" Best Time to Buy and Sell Stock II")
// 	fmt.Println(maxProfit1([]int{7, 1, 5, 3, 6, 4}))
// 	fmt.Println(maxProfit1([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(maxProfit1([]int{7, 6, 4, 3, 1}))

// }
