class Solution {
public:
    int getNWithoutDivizor(int divizor, int n) {
        while (n % divizor == 0) {
            n /= divizor;
        }
        return n;
    }

    bool isUgly(int n) {
        if (n <= 0) return false;
        n = getNWithoutDivizor(2, n);
        n = getNWithoutDivizor(3, n);
        n = getNWithoutDivizor(5, n);
        return n == 1;
    }
};