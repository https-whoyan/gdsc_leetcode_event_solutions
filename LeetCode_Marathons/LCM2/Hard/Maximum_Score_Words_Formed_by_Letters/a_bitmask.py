class Solution:
    def getMask(self, number: int, maskSize: int) -> str:
        mask = ""
        while number != 0:
            mask = str(number % 2) + mask
            number //= 2
        mask = "0" * (maskSize - len(mask)) + mask
        return mask


    def maxScoreWords(self, words: List[str], letters: List[str], score: List[int]) -> int:
        n = len(words)
        MaskCount = 2 ** n

        ans, scoreCounter = 0, 0
        mpLetters = defaultdict(int)

        for myChr in letters: mpLetters[myChr] += 1

        for intMask in range(0, MaskCount):
            mask = self.getMask(intMask, n)

            mpLettersOfMask = defaultdict(int)
            for indexOfMask in range(n):
                if mask[indexOfMask] == "1":
                    for myChr in words[indexOfMask]: mpLettersOfMask[myChr] += 1

            scoreCounter = 0
            for myChr in mpLettersOfMask.keys():
                if mpLettersOfMask[myChr] > mpLetters[myChr]:
                    scoreCounter = 0
                    break
                else: scoreCounter += score[ ord(myChr) - ord('a') ] * mpLettersOfMask[myChr]

            if scoreCounter > ans: ans = scoreCounter

        return ans