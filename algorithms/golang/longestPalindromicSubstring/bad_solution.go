package longestPalindromicSubstring

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	dp := make([][]int, l)
	for i := range dp {
		dp[i] = make([]int, l)
		dp[i][i] = 1
	}

	max := string(s[0])
	for k := 1; k < l; k++ {
		i := 0
		for j := i + k; j < l; j++  {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = j - i + 1
				} else {
					if dp[i+1][j-1] > 0 {
						dp[i][j] = dp[i+1][j-1] + 2
					}
				}

				if dp[i][j] > 0 && j - i > len(max) {
					max = s[i:j+1]
				}
			}
			i++
		}
	}

	return max
}
