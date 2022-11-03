package test

func MainFibonacci(n int) int {

	res := fibonacci(n)
	return res
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
