func canJump(nums []int) bool {
    target := len(nums)-1
    
    
    for i := len(nums)-1; i > -1; i-- {
        if i+nums[i] >= target {
            target = i
        }
    }

    if target == 0 {
        return true
    } else {
        return false
    }
}
