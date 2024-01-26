class Solution {
public:
    int addDigits(int num) {
        while (num >= 10) {
            int copyNum = num;
            int sumOfDigits = 0;
            while (copyNum != 0) {
                sumOfDigits += (copyNum) % 10;
                copyNum /= 10;
            }
            num = sumOfDigits;
        }
        return num;
    }
};