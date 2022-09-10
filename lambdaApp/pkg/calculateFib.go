package pkg

func CalculateFibUsingMemo(index int, memo map[int]int) int {
	//Returns fibonacci number at index provided, by using Memoization
	if val, ok := memo[index]; ok {
		return val
	}
	if index <= 2 {
		return 1
	}
	memo[index] = CalculateFibUsingMemo(index-1, memo) + CalculateFibUsingMemo(index-2, memo)
	return memo[index]
}
