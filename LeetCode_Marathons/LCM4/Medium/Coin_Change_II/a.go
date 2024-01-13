func change(amount int, coins []int) int {
    dp := make([]int, amount + 1)
    dp[0] = 1
    for _, coinAmount := range coins {
        for i := 0; i <= amount; i++ {
            if i - coinAmount >= 0  {
                dp[i] += dp[i - coinAmount]
            }
        }
    }
    return dp[amount]
}