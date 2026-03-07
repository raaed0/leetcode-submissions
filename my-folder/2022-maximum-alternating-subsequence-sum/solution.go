func maxAlternatingSum(nums []int) int64 {
    peak := 0
    valley := 0

    for _, v := range nums {
        peak = max(peak, valley + v)
        valley = max(valley, peak-v)
    }

    return int64(peak)
}
