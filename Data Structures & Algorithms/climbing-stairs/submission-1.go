// func climbStairs(n int) int {
//     if n == 1 || n == 2 {
// 		return n
// 	}

// 	return climbStairs(n - 1) + climbStairs(n-2)
// }

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	a := 0
	b :=1
	for n >= 0{
		temp := a + b
		a = b
		b = temp
		n --
	}
	return a
}