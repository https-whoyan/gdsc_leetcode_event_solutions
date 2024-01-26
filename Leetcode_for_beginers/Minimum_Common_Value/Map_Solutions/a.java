class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        Map<Integer, Boolean> mpNums2 = new HashMap<>();
        for (int number2 : nums2) {
            mpNums2.put(number2, true);
        }

        for (int number1 : nums1) {
            if (mpNums2.containsKey(number1)) {
                return number1;
            }
        }
        return -1;
    }
}