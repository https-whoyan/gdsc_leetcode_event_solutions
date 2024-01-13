class Solution {
    public int[] constructArray(int n, int k) {
        int[] ans = new int[n];
        int index = 0;
        int l = 1, r = k + 1;
        while (l <= r) {
            ans[index++] = l++;
            if (l > r) break;
            ans[index++] = r--;
        }
        for (int i = k + 2; i <= n; i++) {
            ans[index++] = i;
        }
        return ans;
    }
}