class Solution:
    def change(self, amount: int, coins: List[int]) -> int:
        dp: List[int] = [0] * (amount + 1)
        dp[0] = 1
        for coinAmount in coins:
            for i in range(1, amount + 1):
                if i - coinAmount >= 0:
                    dp[i] += dp[i - coinAmount]
        return dp[amount]