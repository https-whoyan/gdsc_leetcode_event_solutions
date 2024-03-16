class Solution {
private:
    bool isMonotic(string word, bool isUpper) {
        for (char letter : word) {
            if (isUpper && !('A' <= letter && letter <= 'Z')) {
                return false;
            }
            if (!isUpper && !('a' <= letter && letter <= 'z')) {
                return false;
            }
        }

        return true;
    }
public:
    bool detectCapitalUse(string word) {
        if ('a' <= word[0] && word[0] <= 'z') {
            return isMonotic(word.substr(1), false);
        }
        return isMonotic(word.substr(1), false) || isMonotic(word.substr(1), true);
    }
};