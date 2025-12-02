func maximumWealth(accounts [][]int) int {
    richest := 0
    for _, m := range accounts {
        sum := 0
        for _, n := range m {
            sum += n
        }
        if sum > richest {
            richest = sum
        }
    }

    return richest
}
