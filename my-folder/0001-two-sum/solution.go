func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				slice := []int{i, j}
				return slice
			}
		}
	}

	return nil
}
