func canAliceWin(nums []int) bool {
    sum_1 := 0
    sum_2 := 0
    for _, v := range nums {
        if v > 9 {
            sum_1 += v
        } else {
            sum_2 += v
        }
    }
    if sum_1 == sum_2 {
        return false
    }
    return true
}
