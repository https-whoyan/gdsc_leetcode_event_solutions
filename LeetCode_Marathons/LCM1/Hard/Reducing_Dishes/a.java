class Solution {
    int[] globalSatisfaction;

    private int getSum(int count) {
        int n = globalSatisfaction.length;
        int ans = 0, counter = 1;
        for (int i = n - count; i <= n - 1; i++) {
            ans += counter * globalSatisfaction[i];
            counter++;
        }

        return ans;
    }

    public int maxSatisfaction(int[] satisfaction) {
        int n = satisfaction.length;
        Arrays.sort(satisfaction);
        globalSatisfaction = satisfaction;
        
        int l = 0, r = n, mid = 0;
        while (l < r) {
            mid = (l + r + 1) / 2;
            if (getSum(mid) > getSum(mid - 1)) l = mid;
            else r = mid - 1;
        }

        return getSum(l);
    }
}