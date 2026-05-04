func reverse(x int) int {
    result := 0
    
    for x != 0 {
        digit := x % 10
        x /= 10

        if result > (math.MaxInt32-digit)/10 || result < (math.MinInt32-digit)/10  {
            return 0
        }

        result = result*10 + digit
    }

    
    return result
}
