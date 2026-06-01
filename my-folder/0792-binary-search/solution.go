func search(nums []int, target int) int {

    var binarySearch func(l, r int) int
    binarySearch = func(l, r int) int {
        if l == r {
            if nums[l] == target {
                return l
            }
            return -1
        }
        mid := l+(r-l)/2
        if target <= nums[mid] {
            return binarySearch(l, mid)
        } else {
            return binarySearch(mid+1,r)
        }
    }

    return binarySearch(0, len(nums)-1)
}
