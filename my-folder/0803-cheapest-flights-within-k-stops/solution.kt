class Solution {
    fun findCheapestPrice(n: Int, flights: Array<IntArray>, src: Int, dst: Int, k: Int): Int {
        val INF = 1_000_000_000
        val adj = Array(n) { mutableListOf<Pair<Int, Int>>() }
        val dist = Array(n) { IntArray(k + 5) { INF } }
        for (flight in flights) {
            adj[flight[0]].add(Pair(flight[1], flight[2]))
        }
        dist[src][0] = 0
        val minHeap = PriorityQueue(compareBy<Triple<Int, Int, Int>> { it.first })
        minHeap.add(Triple(0, src, -1))
        while (minHeap.isNotEmpty()) {
            val (cst, node, stops) = minHeap.poll()
            if (node == dst) return cst
            if (stops == k || dist[node][stops + 1] < cst) continue
            for ((nei, w) in adj[node]) {
                val nextCst = cst + w
                val nextStops = stops + 1
                if (dist[nei][nextStops + 1] > nextCst) {
                    dist[nei][nextStops + 1] = nextCst
                    minHeap.add(Triple(nextCst, nei, nextStops))
                }
            }
        }
        return -1
    }
}
