class Solution:
    def isMonotic(self, word: str, isUpper: bool) -> bool:
        for letter in word:
            if isUpper and not ("A" <= letter and letter <= "Z"):
                return False
            if not isUpper and not ("a" <= letter and letter <= "z"):
                return False
        return True
    

    def detectCapitalUse(self, word: str) -> bool:
        if "a" <= word[0] and word[0] <= "z":
            return self.isMonotic(word[1:], False)
        return self.isMonotic(word[1:], False) or self.isMonotic(word[1:], True)