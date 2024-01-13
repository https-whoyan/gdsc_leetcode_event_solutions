class Solution {
public:
    int search(vector<int>& nums, int target) {
        int n;
        int l, r, mid;
        
        n = nums.size();
        int secondArrayStartIndex = 0;

        if (nums[0] > nums[n - 1]) {
            l = 0, r = n - 1;
            while (l < r) {
                mid = (l + r + 1) / 2;
                if (nums[mid] > nums[l]) l = mid;
                else {
                    if (nums[mid - 1] > nums[mid]) l = mid, r = mid;
                    else r = mid - 1;
                }
            }
            secondArrayStartIndex = l;
        }
        
        l = 0, r = secondArrayStartIndex - 1;
        while (l <= r) {
            mid = (l + r) / 2;
            if (nums[mid] == target) return mid;
            else if (nums[mid] < target) l = mid + 1;
            else r = mid - 1;
        }

        l = secondArrayStartIndex, r = n - 1;
        while (l <= r) {
            mid = (l + r) / 2;
            if (nums[mid] == target) return mid;
            else if (nums[mid] < target) l = mid + 1;
            else r = mid - 1;
        }
        return -1;
    }
};