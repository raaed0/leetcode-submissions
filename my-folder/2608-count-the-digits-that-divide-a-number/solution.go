func countDigits(num int) int {
    tmp := num
    count := 0
    for tmp > 0 {
        d := tmp % 10
        tmp = tmp / 10
        if num % d == 0 {
            count++
        }
    }
    return count
}
