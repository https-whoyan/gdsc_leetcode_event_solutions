class Solution {
public:
    int countPrimes(int n) {
        vector<bool> isNotPrimeStatuses(n + 1, false);
        int counter = 0;
        for (int i = 2; i < n; i++) {
            if (!isNotPrimeStatuses[i]) {
                if (sqrt(n) >= i) {
                    for (int j = i * i; j < n; j += i) isNotPrimeStatuses[j] = true;
                }
                counter++;
            }
        }
        return counter;
    }
};