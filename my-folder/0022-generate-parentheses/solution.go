func generateParenthesis(n int) []string {
    stack := []string{}
    res := []string{}

    var backtrack func(open int, close int)
    backtrack = func(open int, close int) {
        if open == close && close == n {
            res = append(res, strings.Join(stack, ""))
            return
        }

        if open < n {
            stack = append(stack, "(")
            backtrack(open+1, close)
            stack = stack[:len(stack)-1]
        }

        if close < open {
            stack = append(stack, ")")
            backtrack(open, close+1)
            stack = stack[:len(stack)-1]
        }
    }

    backtrack(0, 0)
    return res

}
