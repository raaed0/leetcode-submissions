func numTeams(rating []int) int {
    res := 0
    

    for mid := 1; mid <len(rating)-1; mid++ {
        left_smaller := 0
        left_larger := 0
        right_smaller := 0
        right_larger := 0

        for i := 0; i < mid; i++ {
            if rating[i] < rating[mid] {
                left_smaller++
            }
        }
        for i := mid+1; i<len(rating); i++ {
            if rating[i] > rating[mid] {
                right_larger++
            }
        }

        res += left_smaller * right_larger

        left_larger = mid - left_smaller
        right_smaller = len(rating) - (mid + 1) - right_larger
        
        res += left_larger * right_smaller
    }

    return res
    
}
