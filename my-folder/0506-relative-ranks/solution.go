func findRelativeRanks(score []int) []string {
    n := len(score)

    indexed := make([][2]int, n)
    for i, v := range score {
        indexed[i] = [2]int{v, i}
    }

    sort.Slice(indexed, func(a, b int) bool {  // ← capital S
        return indexed[a][0] > indexed[b][0]
    })

    res := make([]string, n)

    for i, pair := range indexed {
        originalIndex := pair[1]
        switch i {
        case 0:
            res[originalIndex] = "Gold Medal"
        case 1:
            res[originalIndex] = "Silver Medal"
        case 2:
            res[originalIndex] = "Bronze Medal"
        default:
            res[originalIndex] = strconv.Itoa(i + 1)
        }       
    }            

    return res
}
