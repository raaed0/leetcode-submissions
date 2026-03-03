func longestConsecutive(nums []int) int {
    set := make(map[int]bool)
    longest := 0

    for _, n := range nums {
        set[n] = true
    }

    for n := range set {
        if !set[n-1] {
            longer := 1

            for set[n+longer] {
                longer++
            }

            longest = max(longest, longer)
        }
    }

    return longest
}

