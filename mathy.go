package mathy

func Fib(n int) int64 {
	if n <= 1 {
		return 1
	}
	var s = make([]int64, n + 1)
	s[0] = 1
	s[1] = 1
	for i := 2; i <= n ; i++ {
		s[i] = s[i-1] + s[i-2]
	}
	return s[n]
}

func Fact(n int) int64 {
	if n <= 1 {
		return 1
	}
	var s int64 = 1
	for i := 2; i <= n; i++ {
		s *= int64(i)
	}
	return s

}
