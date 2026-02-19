func minimumSum(n int, k int) int {
    res := 0
    
    for i:=1; i<=n; i++ {
        if i <= (k/2) {
            res += i
        } else {
            res += k
            k++
        }
    }

    return res
}
