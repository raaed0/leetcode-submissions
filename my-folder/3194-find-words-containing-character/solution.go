func findWordsContaining(words []string, x byte) []int {
    res := []int{}
    for i, w := range words {
        if strings.IndexByte(w, x) != -1 {
            res = append(res, i)
        }
    }
    return res
}
