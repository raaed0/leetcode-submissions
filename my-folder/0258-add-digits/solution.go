func addDigits(num int) int {
    d := 0
    for num > 0 {
        d += num % 10
        num /= 10
    }
    if d > 9 {
        return addDigits(d)
    }
    return d
}
