class Solution:
    def convertToBase7(self, num: int) -> str:
        if num == 0: return "0"
        needAddMinusOne: bool = num < 0
        if needAddMinusOne: num = -1 * num

        base7s: str = ""
        while num != 0:
            base7s = chr(ord('0') + num % 7) + base7s
            num //= 7

        if needAddMinusOne: base7s = '-' + base7s
        return base7s