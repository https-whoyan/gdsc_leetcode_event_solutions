class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        n: int = len(digits)
        digits[n - 1] += 1
        for i in range(n - 1, -1, -1):
            if digits[i] == 10:
                digits[i] = 0
                if i != 0: digits[i - 1] += 1
                else: return [1] + digits
            else:
                return digits