class Solution {
    public int[][] kClosest(int[][] points, int k) {
        int n = points.length;
        Integer[][] listDistance = new Integer[n][2];
        
        for (int i = 0; i < n; i++) {
            int distance = points[i][0] * points[i][0] + points[i][1] * points[i][1];
            listDistance[i] = new Integer[]{distance, i};
        }
        
        Arrays.sort(listDistance, new Comparator<Integer[]>() {
            public int compare(Integer[] a, Integer[] b) {
                return a[0].compareTo(b[0]);
            }
        });

        int[][] ans = new int[k][2];
        for (int i = 0; i < k; i++) ans[i] = points[listDistance[i][1]];

        return ans;
    }
}