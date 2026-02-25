func findJudge(n int, trust [][]int) int {
    peopleMap := make(map[int]int)

    for _, t := range trust {
        peopleMap[t[0]]--
        peopleMap[t[1]]++
    }

    for p := 1; p<=n; p++ {
        if peopleMap[p] == n-1 {
            return p
        }
    }

    return -1
}
