func countSubstrings(s string) int {
    res := 0

    for i := 0; i < len(s); i++ {
        res += expand(s, i, i)
        res += expand(s, i, i+1)
    }

    return res
}

func expand(str string, left, right int) int {
    count := 0
    for left >=0 && right < len(str) && str[left] == str[right] {
        count++
        left--
        right++
    }
    return count
}
