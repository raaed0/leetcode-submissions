func subtractProductAndSum(n int) int {
    if n == 0 {
        return n
    }

    sum := 0
    product := 1
    for n > 0 {
        d := n%10
        sum += d
        product *= d
        n = n/10
    }
    return product - sum
}
