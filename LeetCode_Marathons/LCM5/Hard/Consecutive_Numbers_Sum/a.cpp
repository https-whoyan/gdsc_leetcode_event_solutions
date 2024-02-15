class Solution {
public:
    vector<int> getAllDivizors(int n) {
        int sqrtN = sqrt(n);
        vector<int> allDivizors;
        for (int i = 1; i <= sqrtN; i++) {
            if (n % i == 0) {
                allDivizors.push_back(i);
                if (n / i != i) allDivizors.push_back(n / i);
            }
        }
        return allDivizors;
    }

    int consecutiveNumbersSum(int n) {
        vector<int> allDivizors = getAllDivizors(n);
        int counter = 0;
        for (int divizor : allDivizors) {
            if (divizor % 2 == 1) counter++;
        }
        return counter;
    }
};