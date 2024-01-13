class Solution {
public:
    int change(int amount, vector<int>& coins) {
        int* dp = new int[amount + 1];
        for (int i = 0; i <= amount; i++) dp[i] = 0;
        dp[0] = 1;
        for (int coinAmount : coins) {
            for (int i = 0; i <= amount; i++) {
                if (i - coinAmount >= 0) {
                    dp[i] += dp[i - coinAmount];
                }
            }
        }
        int ans = dp[amount];
        delete[] dp;
        return ans;
    }
};