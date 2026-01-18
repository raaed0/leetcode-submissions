func subsetsWithDup(nums []int) [][]int {
    sort.Ints(nums)
    var res [][]int
    backtracking(0, []int{}, nums, &res)
    return res
}

func backtracking(i int, path []int, nums []int, res *[][]int) {
    if i == len(nums) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }
    
    backtracking(i+1, append(path, nums[i]), nums, res)

    for i+1 < len(nums) && nums[i] == nums[i+1] {
        i++
    }
    
    // Backtrack
    backtracking(i+1, path, nums, res)
}
