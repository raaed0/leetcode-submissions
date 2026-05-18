class Solution(object):
    def canVisitAllRooms(self, rooms):
        visited = set()
        queue = deque()

        queue.append(0)
        visited.add(0)

        while queue:
            curr = queue.popleft()

            for neighbour in rooms[curr]:
                if neighbour not in visited:
                    visited.add(neighbour)
                    queue.append(neighbour)

        return len(visited) == len(rooms)
        
