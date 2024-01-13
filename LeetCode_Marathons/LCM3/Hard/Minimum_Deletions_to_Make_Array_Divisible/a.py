class Solution:
    def gcd(number1: int, number2: int) -> int:
        while number2 != 0:
            number1, number2 = number2, number1 % number2
        return number1
    

    def minOperations(self, nums: List[int], numsDivide: List[int]) -> int:
        gcdOfNumsDivine: int = numsDivide[0]
        for number in numsDivide:
            gcdOfNumsDivine = gcd(gcdOfNumsDivine, number)
        nums.sort()
        counter: int = 0
        for number in nums:
            if gcdOfNumsDivine % number == 0: return counter
            counter += 1
        return -1