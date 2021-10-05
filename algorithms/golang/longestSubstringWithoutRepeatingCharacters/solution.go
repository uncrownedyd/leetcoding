package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	seq := [128]bool{}

	maxLen := 0
	left := 0
	for i := 0; i < len(s); i++ {
		for seq[s[i]] {
			seq[s[left]] = false
			left++
		}

		if i - left + 1 > maxLen {
			maxLen = i - left + 1
		}
		seq[s[i]] = true
	}
	return maxLen
}
