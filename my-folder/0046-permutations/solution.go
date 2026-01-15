func permute(nums []int) [][]int {
    result := [][]int{}
    used := make([]bool, len(nums))
    currList := make([]int, 0, len(nums))

    var backtrack func()
    backtrack = func() {
        if len(currList) == len(nums) {
            tmp := make([]int, len(currList))
            copy(tmp, currList)
            result = append(result, tmp)
            return
        }
        for i := range nums {
            if used[i] {
                continue
            }
            used[i] = true
            currList = append(currList, nums[i])
            
            backtrack()

            used[i] = false
            currList = currList[:len(currList)-1]
        }
    }

    backtrack()
    return result
}
