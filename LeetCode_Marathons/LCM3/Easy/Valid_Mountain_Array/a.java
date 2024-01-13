class Solution {
    public boolean validMountainArray(int[] arr) {
        int n = arr.length;
        Boolean isIncreasing = true;
        for (int i = 0; i < n - 1; i++) {
            if (isIncreasing) {
                if (arr[i + 1] < arr[i]) {
                    if (i != 0) isIncreasing = false;
                    else return false;
                } else if (arr[i + 1] == arr[i]) return false;
            } else {
                if (arr[i + 1] >= arr[i]) return false;
            }
        }
        return !isIncreasing;
    }
}