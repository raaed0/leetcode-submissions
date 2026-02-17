func wiggleSort(nums []int)  {
    n := len(nums)
    if n == 0 {
        return
    }

    median := findMedian(nums, n/2)

    m := func(idx int) int {
        return (2*idx + 1) % (n | 1)
    }

    first, mid, last := 0, 0, n-1

    for mid <= last {
        if nums[m(mid)] > median {
            nums[m(first)], nums[m(mid)] = nums[m(mid)], nums[m(first)]
            first++
            mid++
        } else if nums[m(mid)] < median {
            nums[m(mid)], nums[m(last)] = nums[m(last)], nums[m(mid)]
            last--
        } else {
            mid++
        }
    }
}

func findMedian(nums []int, k int) int {
    left, right := 0, len(nums)-1
    for left < right {
        pivot := nums[right]
        i := left

        for j :=  left; j < right; j++ {
            if nums[j] < pivot {
                nums[i], nums[j] = nums[j], nums[i]
                i++
            }
        }
        nums[i], nums[right] = nums[right], nums[i]

        if i == k {
            return nums[i]
        } else if i < k {
            left = i+1
        } else {
            right = i-1
        }
    }

    return nums[left]
}
