class Solution:
    def twoEditWords(self, q: List[str], d: List[str]) -> List[str]:
        return [s for s in q if any(deque(takewhile(lambda p:p<4,
            accumulate(map(ne,s,t))),maxlen=1)[0]<3 for t in d)]
