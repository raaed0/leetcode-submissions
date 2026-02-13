func maxSubArray(nums []int) int {
    sum := 0
    maxSub := nums[0]

    for _, n := range nums {
        if sum < 0 {
            sum = 0
        }
        sum += n
        maxSub = max(maxSub, sum)
    }

    return maxSub
}
