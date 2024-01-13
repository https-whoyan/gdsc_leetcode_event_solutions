class Solution {
public:
    vector<string> globalWords;
    vector<int> globalScoreOfWords;
    
    int countAns(int indexWord, map<char, int>& lettersMap) {
        if (indexWord == globalWords.size()) return 0;

        map<char, int> originalState = lettersMap;

        int ansWithoutThisWord = countAns(indexWord + 1, lettersMap);

        for (auto letter : globalWords[indexWord]) {
            if (lettersMap[letter] == 0) {
                lettersMap = originalState;
                return ansWithoutThisWord;
            }
            lettersMap[letter] -= 1;
        }

        int ansWithThisWord = countAns(indexWord + 1, lettersMap) + globalScoreOfWords[indexWord];
        lettersMap = originalState;
        return max(ansWithThisWord, ansWithoutThisWord);
    }

    int maxScoreWords(vector<string>& words, vector<char>& letters, vector<int>& score) {
        globalWords = words;

        map<char, int> mapOfLetters;
        mapOfLetters.clear();
        for (auto letter : letters) mapOfLetters[letter]++;

        for (int i = 0; i < words.size(); i++) {
            globalScoreOfWords.push_back(0);
            for (auto letter : words[i]) globalScoreOfWords[i] += score[letter - 'a'];
        }

        return countAns(0, mapOfLetters);
    }
};