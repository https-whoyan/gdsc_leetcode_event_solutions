class Solution {
    public int smallestRangeI(int[] nums, int k) {
        int minNum = Integer.MAX_VALUE, maxNum = Integer.MIN_VALUE;

        for (Integer num : nums) {
            if (num < minNum) minNum = num;
            if (num > maxNum) maxNum = num;
        }

        if (maxNum - minNum <= 2 * k) return 0;
        return maxNum - minNum - 2 * k;
    }
}