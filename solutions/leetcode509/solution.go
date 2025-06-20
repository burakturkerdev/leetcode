package leetcode509

// unoptimized, calculating same f(x) again and again.
func fib(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
