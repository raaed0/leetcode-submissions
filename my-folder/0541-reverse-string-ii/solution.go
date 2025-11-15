func reverseStr(s string, k int) string {
    b := []byte(s)
    l := 0
    r := min(k, len(b))

    for l < r {
        for i,j := l, r-1; i < j; i,j = i+1, j-1 {
            b[i], b[j] = b[j], b[i]
        }

        l += 2 * k
        r = min(l+k, len(b))
    }

    return string(b)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
