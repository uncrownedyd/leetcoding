package twoSum

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)

	for i, v := range nums {
		dict[v] = i
	}

	for i, v := range nums {
		find := target - v
		if ret, ok := dict[find]; ok {
			if ret == i {
				continue
			}
			return []int{i, ret}
		}
	}
	return []int{}
}
