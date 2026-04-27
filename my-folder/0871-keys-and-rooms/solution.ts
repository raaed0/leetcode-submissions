function canVisitAllRooms(rooms: number[][]): boolean {
    const totalRooms: number = rooms.length;
    const visitedRooms: boolean[] = Array(totalRooms).fill(false);
  
    const exploreRoom = (currentRoom: number): void => {
        if (visitedRooms[currentRoom]) {
            return;
        }
      
        visitedRooms[currentRoom] = true;
    
        for (const nextRoom of rooms[currentRoom]) {
            exploreRoom(nextRoom);
        }
    };
  
    exploreRoom(0);
    return visitedRooms.every((isVisited: boolean) => isVisited);
};
