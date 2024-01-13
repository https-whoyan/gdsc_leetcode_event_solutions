class Solution:
    globalWords = []
    globalScoreOfWords = []

    def countAns(self, indexWord: int, lettersMap: dict()) -> int:
        if indexWord == len(self.globalWords): return 0

        ansWithoutThisWord = self.countAns(indexWord + 1, lettersMap.copy())

        for letter in self.globalWords[indexWord]:
            if lettersMap[letter] == 0: return ansWithoutThisWord
            lettersMap[letter] -= 1

        ansWithThisWord = self.countAns(indexWord + 1, lettersMap) + self.globalScoreOfWords[indexWord]
        return max(ansWithoutThisWord, ansWithThisWord)
    

    def maxScoreWords(self, words: List[str], letters: List[str], score: List[int]) -> int:
        self.globalWords = words
        self.globalScoreOfWords = [0] * len(words)

        mapOfLetters = defaultdict(int)
        for letter in letters: mapOfLetters[letter] += 1

        for i in range(len(words)): 
            for letter in words[i]: self.globalScoreOfWords[i] += score[ ord(letter) - ord('a') ]

        return self.countAns(0, mapOfLetters)