func checkIfPangram(sentence string) bool {
    if len(sentence) < 26 {
        return false
    }
    
    for i := 97; i < 123; i++ {
        if !strings.Contains(sentence, string(byte(i))) {
            return false
        }
    }
    return true
}
