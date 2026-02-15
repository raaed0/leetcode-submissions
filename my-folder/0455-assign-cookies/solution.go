func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
    childIdx, cookieIdx := 0, 0

    for childIdx < len(g) && cookieIdx < len(s) {
        if s[cookieIdx] >= g[childIdx] {
            childIdx++
        }
        cookieIdx++
    }

    return childIdx
}
