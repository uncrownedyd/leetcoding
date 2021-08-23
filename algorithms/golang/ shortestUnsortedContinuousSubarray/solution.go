package shortestUnsortedContinuousSubarray

import "math"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	min, max := math.MaxInt64, math.MinInt64

	var left, right int
	for i := n - 1; i >= 0; i-- {
		if nums[i] <= min {
			min = nums[i]
		} else {
			left = i
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] >= max {
			max = nums[i]
		} else {
			right = i
		}
	}

	if right<=left{
		return 0
	}

	return right-left+1

}