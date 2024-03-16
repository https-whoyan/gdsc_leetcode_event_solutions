class Solution {
public:
    int countHomogenous(string s) {
        long long mod = 1000000007, ans = 0, keyCounter = 0;
        char key = '*';
        s += "*";

        for (char myCh : s) {
            if (myCh == key) keyCounter++;
            else {
                long long ansForKey;
                ansForKey = (keyCounter + 1) * keyCounter / 2;
                ans = (ans + ansForKey) % mod;
                keyCounter = 1;
                key = myCh;
            }
        }

        return ans;
    }
};