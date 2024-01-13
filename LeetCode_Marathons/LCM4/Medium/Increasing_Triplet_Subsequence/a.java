class Solution {
    public boolean increasingTriplet(int[] nums) {
        int value1 = 2147483647, value2 = 2147483647;
        for (int num : nums) {
            if (num <= value1) value1 = num;
            else if (num <= value2) value2 = num;
            else return true;
        }
        return false;
    }
}