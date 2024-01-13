class Solution {
public:
    string getMask(int number, int maskSize) {
        string mask;
        while (number != 0) {
            mask = char( int('0') + (number % 2)) + mask;
            number /= 2;
        }
        int needToAppend = maskSize - mask.size();
        for (int jj = 0; jj < needToAppend; jj++) mask = "0" + mask;
        return mask;
    }

    int maxScoreWords(vector<string>& words, vector<char>& letters, vector<int>& score) {
        int n = words.size();
        int maskCount = pow(2, n);

        int ans = 0, scoreCounter = 0;
        map<char, int> mpLetters;

        for (auto myChr : letters) mpLetters[myChr]++;

        for (int maskInt = 1; maskInt < maskCount; maskInt++) {
            string mask = getMask(maskInt, n);
            map<char, int> mpLettersOfMask;

            for (int indexOfMask = 0; indexOfMask < n; indexOfMask++) {
                if (mask[indexOfMask] == '1') {
                    for (auto myChr : words[indexOfMask]) mpLettersOfMask[myChr]++;
                }
            }

            scoreCounter = 0;
            for (auto myPair : mpLettersOfMask) {
                char key = myPair.first;
                int value = myPair.second, value2 = mpLetters[key];
                if (value > value2) {
                    scoreCounter = 0;
                    break;
                } else {
                    scoreCounter += score[ int(key) - int('a') ] * value;
                }
            }

            if (scoreCounter > ans) ans = scoreCounter;
            mpLettersOfMask.clear();
        }
        
        return ans;
    }
};