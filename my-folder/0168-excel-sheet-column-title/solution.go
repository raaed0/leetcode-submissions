
func convertToTitle(columnNumber int) string {
    if columnNumber == 0 {
        return ""
    }
    n := columnNumber - 1
    return convertToTitle(n/26) + string(rune('A'+n%26))
}
