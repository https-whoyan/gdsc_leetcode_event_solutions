class Solution {
public:
    string convertToBase7(int num) {
        if (num == 0) return "0";
        bool needAddMinusOne = num < 0;
        if (needAddMinusOne) num = -1 * num;

        string base7s;
        while (num != 0) {
            base7s = char('0' + num % 7) + base7s;
            num /= 7;
        }

        if (needAddMinusOne) base7s = '-' + base7s;
        return base7s;
    }
};