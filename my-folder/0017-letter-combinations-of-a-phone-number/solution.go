func letterCombinations(digits string) []string {
    res := []string{}
    if len(digits) == 0 {
        return res
    }

    letters := map[byte]string {
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

    var backtracking func(i int, currString string)
    backtracking = func(i int, currString string) {
        if len(currString) == len(digits) {
            res = append(res, currString)
            return
        }

        for _, c := range letters[digits[i]] {
            backtracking(i+1, currString+string(c))
        }
    }

    backtracking(0, "")
    return res
}
