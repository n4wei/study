package fib

// 5:16-5:31

// runtime O(2^n)
// space O(n)
func fibRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

/*
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
*/

// runtime O(n*n)
// space O(n)
func fibRecursiveMemo(n int) int {
	if n < 2 {
		return n
	}
	cache := map[int]int{}
	cache[0] = 0
	cache[1] = 1
	return fibMemo(cache, n)
}

func fibMemo(cache map[int]int, k int) int {
	val, exist := cache[k]
	if !exist {
		cache[k] = fibMemo(cache, k-1) + fibMemo(cache, k-2)
		val = cache[k]
	}
	return val
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