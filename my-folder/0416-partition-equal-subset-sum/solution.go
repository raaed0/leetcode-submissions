func canPartition(nums []int) bool {
    target := 0
    for _, v := range nums {
        target += v
    }

    if target % 2 == 1 {
        return false
    }

    target = target / 2
    
    set := make(map[int]struct{})
    set[0] = struct{}{}
    for i := len(nums)-1; i >= 0; i-- {
        nextSet := make(map[int]struct{})
        for v := range set {
            if v == target || v + nums[i] == target {
                return true
            }
            nextSet[v] = struct{}{}
            nextSet[v + nums[i]] = struct{}{}
        }
        set = nextSet
    }

    _, ok := set[target]

    return ok
}
