class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>>& points, int k) {
        int n = points.size();
        vector<vector<int>> listDistance;

        for (int i = 0; i < n; i++) {
            int distance = points[i][0] * points[i][0] + points[i][1] * points[i][1];
            listDistance.push_back({distance, i});
        }

        sort(listDistance.begin(), listDistance.end());

        vector<vector<int>> ans;
        for (int i = 0; i < k; i++) ans.push_back(points[listDistance[i][1]]);

        return ans;
    }
};