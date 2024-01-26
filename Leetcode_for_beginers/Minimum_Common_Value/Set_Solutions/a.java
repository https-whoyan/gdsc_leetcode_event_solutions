class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        Set<Integer> s1 = new HashSet<>();
        Set<Integer> s2 = new HashSet<>();
        for (int number1 : nums1) {
            s1.add(number1);
        }
        for (int number2 : nums2) {
            s2.add(number2);
        }

        s1.retainAll(s2);
        if (s1.size() == 0) return -1;
        int minEl = Integer.MAX_VALUE;
        for (int commonEl : s1) {
            if (commonEl < minEl) minEl = commonEl;
        }
        return minEl;
    }
}