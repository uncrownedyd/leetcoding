package threeSum

import "sort"

func threeSum(nums []int) [][]int {
	var ret [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums) - 2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		cur := nums[i]

		l := i + 1
		r := len(nums) - 1

		for l < r {
			if cur + nums[l] + nums[r] > 0 {
				r--
			} else if cur + nums[l] + nums[r] < 0 {
				l++
			} else {
				ret = append(ret, []int{cur, nums[l], nums[r]})

				for l < r && l + 1 < len(nums) && nums[l] == nums[l+1] {
					l++
				}

				for r > l && r - 1 > 0 && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			}
		}
	}
	return ret
}