package main

import (
	"fmt"
)

// 5:16-5:31
func main() {
	fmt.Println("fibonacci recursive")
	fmt.Println(fibRecursive(0))
	fmt.Println(fibRecursive(1))
	fmt.Println(fibRecursive(2))
	fmt.Println(fibRecursive(3))
	fmt.Println(fibRecursive(4))
	fmt.Println(fibRecursive(5))
	fmt.Println(fibRecursive(6))

	fmt.Println("fibonacci recursive with memoization")
	fmt.Println(fibRecursiveMemo(0))
	fmt.Println(fibRecursiveMemo(1))
	fmt.Println(fibRecursiveMemo(2))
	fmt.Println(fibRecursiveMemo(3))
	fmt.Println(fibRecursiveMemo(4))
	fmt.Println(fibRecursiveMemo(5))
	fmt.Println(fibRecursiveMemo(6))

	fmt.Println("fibonacci iterative")
	fmt.Println(fibIterative(0))
	fmt.Println(fibIterative(1))
	fmt.Println(fibIterative(2))
	fmt.Println(fibIterative(3))
	fmt.Println(fibIterative(4))
	fmt.Println(fibIterative(5))
	fmt.Println(fibIterative(6))
}

// runtime O(2^n)
// space O(n)
func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// runtime O(n*n)
// space O(n)
func fibRecursiveMemo(n int) int {
	if n < 2 {
		return n
	}
	memo := make([]int, n+1, n+1)
	memo[1] = 1
	return fibMemo(memo, n)
}

func fibMemo(memo []int, n int) int {
	if n == 0 || memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibMemo(memo, n-1) + fibMemo(memo, n-2)
	return memo[n]
}

// runtime O(n)
// space O(n)
func fibIterative(n int) int {
	if n < 2 {
		return n
	}

	fib := make([]int, n+1, n+1)
	fib[0] = 0
	fib[1] = 1
	for i:=2; i<=n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib[n]
}