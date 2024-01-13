class Solution:
    def countPrimes(self, n: int) -> int:
        isNotPrimeStatuses = [False] * n 
        counter = 0

        for i in range(2, n):
            if not isNotPrimeStatuses[i]:
                if n ** 0.5 >= i:
                    for j in range(i * i, n, i): isNotPrimeStatuses[j] = True
                counter += 1

        return counter