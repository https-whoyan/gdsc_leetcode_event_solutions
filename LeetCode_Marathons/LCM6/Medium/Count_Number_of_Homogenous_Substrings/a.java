class Solution {
    public int countHomogenous(String s) {
        long mod = 1000000007;
        long ans = 0;
        Character key = '*';
        s += "*";
        long keyCounter = 0;
        for (Character myCh : s.toCharArray()) {
            if (myCh == key) keyCounter++;
            else {
                long ansForKey;
                ansForKey = (keyCounter + 1) * keyCounter / 2;
                ans = (ans + ansForKey) % mod;
                keyCounter = 1;
                key = myCh;
            }
        }

        return (int) ans;
    }
}