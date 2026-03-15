func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    res, curr, prev := 0, 0, 0

    for i:=0; i<n-1; i++ {
        prev, curr = curr, max(curr, prev+nums[i])
    }
    res = curr
    prev, curr = 0, 0
    for i:=1; i<n; i++ {
        prev, curr = curr, max(curr, prev + nums[i])
    }

    res = max(res, curr)
    return res
}
