func toLowerCase(s string) string {
    byteSlice := []byte(s)
    for i, ch := range byteSlice {
        if ch >= 'A' && ch <= 'Z' {
            byteSlice[i] = ch + 32
        }
    }
    return string(byteSlice)
}
