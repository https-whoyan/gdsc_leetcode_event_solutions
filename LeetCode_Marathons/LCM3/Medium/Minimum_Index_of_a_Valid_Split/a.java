class Solution {
    private int getdDominantEl(List<Integer> arr) {
        List<Integer> copyArr = new ArrayList<>(arr);
        Collections.sort(copyArr);
        return copyArr.get(arr.size() / 2);
    }

    public int minimumIndex(List<Integer> nums) {
        int domEl = getdDominantEl(nums);
        int n = nums.size();
        int[] pref = new int[n];
        int[] postf = new int[n];
        for (int i = 0; i <= n - 1; i++) {
            pref[i] = 0;
            if (i != 0) pref[i] += pref[i - 1];
            if (nums.get(i) == domEl) pref[i]++;
        }
        for (int i = n - 1; i >= 0; i--) {
            postf[i] = 0;
            if (i != n - 1) postf[i] += postf[i + 1];
            if (nums.get(i) == domEl) postf[i]++;
        }
        for (int i = 0; i < n - 1; i++) {
            int lenFMsv = i + 1;
            int lenSMsv = n - lenFMsv;
            if (pref[i] * 2 > lenFMsv && postf[i + 1] * 2 > lenSMsv) {
                return i;
            }
        }
        return -1;
    }
}