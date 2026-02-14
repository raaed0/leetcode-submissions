func canCompleteCircuit(gas []int, cost []int) int {
    if sum(gas) < sum(cost) {
        return -1
    }

    total := 0
    idx := 0
    for i, _ := range gas {
        total += gas[i] - cost[i]

        if total < 0 {
            total = 0
            idx = i+1
        }
    }

    return idx
}

func sum(slice []int) int {
    total := 0
    for _, n := range slice {
        total += n
    }

    return total
}
