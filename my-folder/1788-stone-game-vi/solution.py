class Solution:
    def stoneGameVI(self, aliceValues: List[int], bobValues: List[int]) -> int:
        heap = []
        n = len(aliceValues)
        for i in range(n):
            heapq.heappush(heap, (-(aliceValues[i] + bobValues[i]), aliceValues[i], bobValues[i]))

        is_alice = 1  # 1=Alice, 0=Bob
        alice = bob = 0
        while heap:
            total, a_val, b_val = heapq.heappop(heap)
            if is_alice:
                alice += a_val

            else:
                bob += b_val
            is_alice = 1 - is_alice

        if alice > bob:
            return 1
        elif alice < bob:
            return -1
        else:
            return 0
