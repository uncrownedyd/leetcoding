package containerWithMostWater

import "math"

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	res := math.MinInt32

	for ; l < r ; {
		res = max(res, min(height[l], height[r]) * (r - l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}