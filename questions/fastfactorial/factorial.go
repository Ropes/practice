package fastfactorial

func FactFast(n int) int {

	return 0
}

func FastFib(n int) int {
	f0, f1 := 0, 1
	if n == 0 {
		return f0
	}
	if n < 2 {
		return f1
	}

	for i := 0; i < n; i++ {
		tmp := f0 + f1
		f0, f1 = f1, tmp
	}

	return f0
}

func SlowFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return SlowFib(n-1) + SlowFib(n-2)
}
