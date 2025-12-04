func differenceOfSum(nums []int) int {
    sum := 0
    digit_sum := 0
    for _, v := range nums {
        sum += v

        for v > 0 {
            digit_sum += v % 10
            v /= 10
        }
    }

    diff := sum - digit_sum
    if diff < 0 {
        diff = -diff
    }
    return diff
}
