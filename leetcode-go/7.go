package main

func reverse(x int) int {
	sign := 1

	// Handle negatives
	if x < 0 {
		sign = -1
		x = x * -1
	}

	num := 0
	for x > 0 {
		remainder := x % 10
		num *= 10
		num += remainder
		x /= 10
	}

	num = sign * num
	// Could use math.MaxInt32 here as well actually
	if num < -2147483648 || num > 2147483647 {
		num = 0
	}

	return num
}
