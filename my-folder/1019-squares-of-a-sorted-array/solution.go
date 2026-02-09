func sortedSquares(nums []int) []int {
    res := []int{}

    for _, num := range nums {
        res = append(res, num * num)
    }

    sort.Ints(res)
    return res
}
