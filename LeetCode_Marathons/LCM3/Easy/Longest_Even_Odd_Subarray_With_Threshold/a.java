class Solution {
    public int longestAlternatingSubarray(int[] nums, int threshold) {
        int n = nums.length;
        for (int arrSz = n; arrSz >= 1; arrSz--) {
            for (int startArr = 0; startArr <= n - arrSz; startArr++) {
                Boolean isCorrect = true;
                if (nums[startArr] % 2 != 0) {
                    isCorrect = false;
                    continue;
                }
                for (int i = startArr; i <= arrSz + startArr - 2; i++) {
                    if (nums[i] % 2 == nums[i + 1] % 2 || nums[i] > threshold) {
                        isCorrect = false;
                        break;
                    }
                }
                if (nums[arrSz + startArr - 1] > threshold) {
                    isCorrect = false;
                }
                if (isCorrect) return arrSz;
            }
        }
        return 0;
    }
}