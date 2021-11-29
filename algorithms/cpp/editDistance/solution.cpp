class Solution {
public:
    int minDistance(string word1, string word2) {
        int l1 = word1.size();
        int l2 = word2.size();
        vector<int> dp(l1 + 1);

        for (int i = 1; i < l1+1; ++i) {
            dp[i] = i
        }

        for (int i = 1; i < l2 + 1; ++i) {
            int pre = i - 1;
            dp[0] = i;
            for (int j = 1; j < l1 + 1; ++j) {
                int tmp = dp[j];
                if (word1[j-1] == word2[i-1]) {
                    dp[j] = pre;
                } else {
                    dp[j] = min(pre, min(dp[j-1], dp[j])) + 1;
                }
                pre = tmp;
            }
        }
        return dp[l1]
    }

    int min(int a, b) {
        return b < a ? b : a;
    }
};