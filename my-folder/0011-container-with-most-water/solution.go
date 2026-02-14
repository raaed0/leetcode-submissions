func maxArea(height []int) int {
    i := 0
    j := len(height) - 1
    maxarea := 0

    for i < j {

        area := min(height[i], height[j]) * (j-i)
        maxarea = max(maxarea, area)
        if height[i] > height[j] {
            j--
        } else {
            i++
        }
    }

    return maxarea
}
