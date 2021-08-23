package getMaximumInGeneratedArray

//You are given an integer n. An array nums of length n + 1 is generated in the following way:
//
//nums[0] = 0
//nums[1] = 1
//nums[2 * i] = nums[i] when 2 <= 2 * i <= n
//nums[2 * i + 1] = nums[i] + nums[i + 1] when 2 <= 2 * i + 1 <= n
//Return the maximum integer in the array nums

func getMaximumGenerated(n int) (ans int) {
	if n == 0 {
		return
	}
	nums := make([]int, n+1)
	nums[1] = 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i/2] + i%2*nums[i/2+1]
	}
	for _, v := range nums {
		ans = max(ans, v)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
