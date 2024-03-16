class Solution {
    public int minPatches(int[] nums, int n) {
        long dpSum = 0;
        int counter = 0;

        ArrayList<Long> arr = new ArrayList<>();
        for (int num: nums) arr.add((long) num);
        arr.add(((long) n) + 1);

        for (long num: arr) {
            while (dpSum + 1 < num) {
                if (dpSum >= (long) n) return counter;
                counter++;
                dpSum = dpSum + 1 + dpSum;
            }
            dpSum += num;
            if (dpSum >= (long) n) return counter;
        }

        return counter;
    }
}