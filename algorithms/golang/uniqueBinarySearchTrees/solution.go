package uniqueBinarySearchTrees

func numTrees(n int) int {
	return helper(1, n)
}

func helper(start, end int) int {
	if start >= end {
		return 1
	}
	if end-start == 1 {
		return 2
	}
	cnt := 0
	for idx := start; idx <= end; idx++ {
		cnt += helper(start, idx-1) * helper(idx+1, end)
	}
	return cnt
}
