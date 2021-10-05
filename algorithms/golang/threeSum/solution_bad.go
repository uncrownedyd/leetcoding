package threeSum

import (
	"math"
	"sort"
)

func threeSumBadSolution(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, 3)

	var last int
	for i, v := range nums {
		if i > 0 && v == last {
			continue
		}
		last = v
		findTar := 0 - v
		sub := findTwo(nums[i+1:], findTar)
		for _, p := range sub {
			res = append(res, append(p, v))
		}
	}
	return res
}

func findTwo(nums []int, target int) [][]int {
	res := make([][]int, 0, 2)
	var last int
	for i, v := range nums {
		if i > 0 && v == last {
			continue
		}
		last = v

		findTar := target - v
		if r := binaryFind(nums[i+1:], findTar); r > math.MinInt32 {
			res = append(res, []int{v, r})
		}
	}
	return res
}

func binaryFind(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			return nums[mid]
		}
	}
	return math.MinInt32
}