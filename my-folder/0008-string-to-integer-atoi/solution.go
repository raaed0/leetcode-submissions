func myAtoi(s string) int {
    i := 0
    l := len(s)
    sign := 1
    for i < l && s[i] == ' ' {
        i++
    }

    if i<l && (s[i] == '-' || s[i] == '+') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }
    result := 0
    for i < l && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')

        if result > (math.MaxInt32 - digit) / 10 {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }

        result = result * 10 + digit
        i++
    }

    return sign * result

}
