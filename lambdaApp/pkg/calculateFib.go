package pkg

func CalculateFibUsingMemo(number int, memo map[int]int) int {
	//calculates fibonachi number at number index provided with Memoization
	if val, ok := memo[number]; ok {
		return val
	}
	if number <= 2 {
		return 1
	}
	memo[number] = CalculateFibUsingMemo(number-1, memo) + CalculateFibUsingMemo(number-2, memo)
	return memo[number]
}
