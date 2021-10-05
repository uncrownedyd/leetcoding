package frogStrepToHome

var rightMost int
func canGetHome(a, x, y int, trap []int) int {
	rightMost = a + x * y

	return helper(a, x, y, trap, []int{})
}

//
//func dfs(a, x, y int, trap []int) int {
//	dp := make([]int, 0, a + x * y)
//	for _, p := range trap {
//		dp[p] = -1
//	}
//}

func helper(a, x, y int ,trap []int, visited []int) int {
	if a < 0 || contains(trap, a) || a > rightMost || contains(visited, a) {
		return -1
	}

	if a == 0 {
		return 1
	}

	visited = append(visited, a)

	fs := helper(a - x, x, y, trap, visited)
	bs := helper(a + y, x, y, trap, visited)
	minPos := minPositive(fs, bs)
	if minPos < 0 {
		return -1
	}
	return minPos + 1
}

func minPositive(a, b int) int {
	if a < 0 && b < 0 {
		return -1
	}
	if a > b {
		return a
	}
	return b
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}