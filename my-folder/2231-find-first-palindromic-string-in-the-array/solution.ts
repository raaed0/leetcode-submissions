function firstPalindrome(words: string[]): string {
    for (let w of words) {
        if (w === w.split("").reverse().join("")) {
            return w
        }
    }
    return ""
};
