class Solution:
    def addDigits(self, num: int) -> int:
        while num >= 10:
            copyNum: int = num
            sumOfDigits: int = 0
            while copyNum != 0:
                sumOfDigits += copyNum % 10
                copyNum //= 10
            num = sumOfDigits
        
        return num