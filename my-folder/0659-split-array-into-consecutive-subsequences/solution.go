func isPossible(nums []int) bool {
    count := make(map[int]int)
    sequenceTail := make(map[int]int)

    for _, n := range nums {
        count[n]++
    }

    for _, n := range nums {
        if count[n] <= 0 {
            continue
        }

        if sequenceTail[n-1] > 0 {
            sequenceTail[n-1]--
            sequenceTail[n]++
            count[n]--
        } else if count[n+1] > 0 && count[n+2] > 0 {
            sequenceTail[n+2]++
            count[n]--
            count[n+1]--
            count[n+2]--
        } else {
            return false
        }
    }

    return true
}
