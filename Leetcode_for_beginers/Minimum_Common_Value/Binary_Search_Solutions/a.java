class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        int n2 = nums2.length;
        for (int number1 : nums1) {
            int l = 0, r = n2 - 1, mid = 0;
            while (l < r) {
                mid = (l + r) / 2;
                if (nums2[mid] == number1) return number1;
                else if (nums2[mid] < number1) l = mid + 1;
                else r = mid - 1;
            }
            if (nums2[l] == number1) return number1;
        }
        return -1;
    }
}