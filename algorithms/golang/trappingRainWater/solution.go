package trappingRainWater

func trap(height []int) int {
	left, right := 0, len(height)-1
	lMax, rMax := height[left], height[right]
	ans := 0

	for left < right {
		if lMax < rMax {
			ans += lMax - height[left]
			left++
			lMax = max(lMax, height[left])
		} else {
			ans += rMax - height[right]
			right--
			rMax = max(rMax, height[right])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
