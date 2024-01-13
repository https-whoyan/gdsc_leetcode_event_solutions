class Solution {
public:
    int peakIndexInMountainArray(vector<int>& arr) {
        int l = 1, r = arr.size() - 2, mid = 0;
        while (l < r) {
            mid = (l + r) / 2;
            if (arr[mid - 1] < arr[mid]) {
                if (arr[mid] < arr[mid + 1]) l = mid + 1;
                else return mid;
            } else r = mid - 1;
        }
        return l;
    }
};