func firstPalindrome(words []string) string {
    for _, w := range words {
        runes := []rune(w)
        i, j := 0, len(runes)-1
        isPalindrome := true
        for i < j {
            if runes[i] != runes[j] {
                isPalindrome = false
                break
            }
            i++
            j--
        }
        if isPalindrome {
            return w
        }
    }
    return ""
}


