class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        vector<int> ansMsv;
        int id1 = 0, id2 = 0;

        while (id1 + id2 != n + m) {
            if (id1 == m) {
                ansMsv.push_back(nums2[id2]);
                id2++;
            } else if (id2 == n || nums1[id1] < nums2[id2]) {
                ansMsv.push_back(nums1[id1]);
                id1++;
            } else {
                ansMsv.push_back(nums2[id2]);
                id2++;
            }
        }

        for (int i = 0; i < n + m; i++) nums1[i] = ansMsv[i];
    }
};