# Algo-master-leetcode-array

### Code
```go
func productExceptSelf(nums []int) []int {
    prefix := 1
    var result []int = make([]int, len(nums))
    for index, value := range nums {
        result[index] = prefix
        prefix = prefix * value
    }
    postfix := 1
    for i := len(nums) - 1; i >= 0; i-- {
        result[i] = result[i] * postfix
        postfix = postfix * nums[i]
    }
    return result
}
```



### Code

```go
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
```