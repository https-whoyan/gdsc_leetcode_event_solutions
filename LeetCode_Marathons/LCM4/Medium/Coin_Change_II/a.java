class Solution {
    public int change(int amount, int[] coins) {
        int[] dp = new int[amount + 1];
        dp[0] = 1;
        for (int coinAmount : coins) {
            for (int i = 0; i <= amount; i++) {
                if (i - coinAmount >= 0) {
                    dp[i] += dp[i - coinAmount];
                }
            }
        }
        return dp[amount];
    }
}