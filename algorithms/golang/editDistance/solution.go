package editDistance

func minDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	dp := make([][]int, l2 + 1)

	for i := range dp {
		dp[i] = make([]int, l1+1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := 1; i < l2 + 1; i++ {
		for j := 1; j < l1 + 1; j++ {
			if word1[j-1] == word2[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))+1
			}
		}
	}

	return dp[l2][l1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}