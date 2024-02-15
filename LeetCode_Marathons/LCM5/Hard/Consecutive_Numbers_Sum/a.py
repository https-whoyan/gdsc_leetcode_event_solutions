class Solution:
    def getAllDivizors(self, n: int) -> List[int]:
        allDivizors: List[int] = []
        sqrtN: int = int(n ** 0.5)
        for i in range(1, sqrtN + 1):
            if n % i == 0:
                allDivizors.append(i)
                if n / i != i: allDivizors.append(int(n / i))
        return allDivizors

    def consecutiveNumbersSum(self, n: int) -> int:
        allDivizors: List[int] = self.getAllDivizors(n)
        counter: int = 0
        for divizor in allDivizors:
            if divizor % 2 == 1: counter += 1
        return counter