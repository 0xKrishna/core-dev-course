package maps

import "fmt"

// Fibonacci (Fibonacci(n)) with memoization

func Map_4(n int) int {
	fmt.Println("Fibonacci (Fibonacci(n)) with memoization")
	fmt.Println("Map_4 - N:", n)

	memo := make(map[int]int)

	result := fibHelper(n, memo)

	fmt.Println("Map_4 - Result:", result)

	return result
}

func fibHelper(n int, memo map[int]int) int {
	if n == 0 || n == 1 {
		return n
	}

	if _, ok := memo[n]; !ok {
		memo[n] = fibHelper(n-1, memo) + fibHelper(n-2, memo)
	}

	return memo[n]
}
