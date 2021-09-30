package climbingStairs

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	last := 1
	cur := 2
	var tmp int
	for i := 2; i < n; i++ {
		tmp = last + cur
		last = cur
		cur = tmp
	}

	return cur
}
