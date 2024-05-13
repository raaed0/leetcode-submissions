func twoSum(nums []int, target int) []int {
    length := len(nums)
    for i := 0; i < length-1; i++ {
		for j := i+1; j < length; j++ {
			if nums[i]+nums[j] == target {
				slice := []int{i, j}
				return slice
			}
		}
	}

	return []int{}
}
