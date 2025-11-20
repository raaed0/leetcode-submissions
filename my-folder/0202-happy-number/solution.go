func isHappy(n int) bool {
    past := make(map[int]bool)

    for n != 1 && !past[n] {
        past[n] = true
        sum := 0
        for n > 0 {
            digit := n % 10
            sum += digit * digit
            n /= 10
        }
        n = sum
    }
    return n == 1
}
