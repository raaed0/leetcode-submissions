class Solution:
    def kthSmallestPrimeFraction(self, arr: List[int], k: int) -> List[int]:
        return nsmallest(k,product(arr,arr),lambda q:q[0]/q[1])[-1]
