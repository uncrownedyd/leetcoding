package medianOfTwoSortedArrays

func findMedianSortedArraysSolution1(nums1 []int, nums2 []int) float64 {
	var aIdx, bIdx, left, right int
	len1 := len(nums1)
	len2 := len(nums2)
	k := len1 + len2
	for i := 0; i <= k / 2; i++ {
		left = right
		if aIdx < len1 && (bIdx >= len2 || nums1[aIdx] < nums2[bIdx]) {
			right = nums1[aIdx]
			aIdx++
		} else {
			right = nums2[bIdx]
			bIdx++
		}
	}

	if k % 2 == 0 {
		return (float64(left) + float64(right)) / 2.0
	} else {
		return float64(right)
	}
}
