class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        n = len(points)
        listDistance = []
        for i in range(n): listDistance.append([points[i][0] ** 2 + points[i][1] ** 2, i])
        listDistance.sort()

        return [points[listDistance[i][1]] for i in range(k)]