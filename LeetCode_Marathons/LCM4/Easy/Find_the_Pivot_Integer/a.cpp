class Solution {
public:
    int pivotInteger(int n) {
        int sumN = n * (n + 1) / 2;
        int dpSum = 0;
        for (int i = 0; i <= n; i++) {
            dpSum += i;
            if (dpSum == sumN - dpSum + i) return i;
        }
        return -1;
    }
};