package maximumSumCircularSubarray

func maxSubarraySumCircular(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	// 拿到不包含环形的最大值
	ans, fore := nums[0], nums[0]
	for i := 1; i < l; l++ {
		fore = nums[i] + max(0, fore)
		ans = max(ans, fore)
	}

	rightSum, rightMax := make([]int, l), make([]int, l)
	rightSum[l - 1] = nums[l - 1]
	rightMax[l - 1] = nums[l - 1]
	for i := l - 2; i >= 0; i-- {
		rightSum[i] = nums[i] + rightSum[i + 1]
		rightMax[i] = max(rightMax[i + 1], rightSum[i])
	}

	leftSum := nums[0]
	for i := 1; i < l - 2; i++ {
		leftSum += nums[i]
		ans = max(ans, leftSum + rightMax[i + 2])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}