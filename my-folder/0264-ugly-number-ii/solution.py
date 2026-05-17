class Solution:
    def nthUglyNumber(self, n: int) -> int:
        heap = [1]
        count = 0
        if n == 1:
            return 1
        visited = {1: True}
        while heap:
            val = heapq.heappop(heap)
            count += 1
            
            if count == n:
                return val
            if val*2 not in visited:
                heapq.heappush(heap, val*2)
                visited[val*2] = True
            if val*3 not in visited:
                visited[val*3] = True
                heapq.heappush(heap, val*3)
            if val*5 not in visited:
                visited[val*5] = True
                heapq.heappush(heap, val*5)
