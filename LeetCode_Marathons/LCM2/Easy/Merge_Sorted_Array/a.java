class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        ArrayList<Integer> ansMsv = new ArrayList<Integer>();
        int id1 = 0, id2 = 0;

        while (id1 + id2 != n + m) {
            if (id1 == m) {
                ansMsv.add(nums2[id2]);
                id2++;
            } else if (id2 == n || nums1[id1] < nums2[id2]) {
                ansMsv.add(nums1[id1]);
                id1++;
            } else {
                ansMsv.add(nums2[id2]);
                id2++;
            }
        }

        for (int i = 0; i < n + m; i++) nums1[i] = ansMsv.get(i);
    }
}