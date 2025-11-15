func evalRPN(tokens []string) int {
    stack := []int{}

    for _, token := range tokens {
        if token == "+" {
            l := len(stack)
            v1, v2 := stack[l-2], stack[l-1]
            stack = stack[:l-2]
            value := v1 + v2
            stack = append(stack, value)
        } else if token == "-" {
            l := len(stack)
            v1, v2 := stack[l-2], stack[l-1]
            stack = stack[:l-2]
            value := v1 - v2
            stack = append(stack, value)
        } else if token == "*" {
            l := len(stack)
            v1, v2 := stack[l-2], stack[l-1]
            stack = stack[:l-2]
            value := v1 * v2
            stack = append(stack, value)
        } else if token == "/" {
            l := len(stack)
            v1, v2 := stack[l-2], stack[l-1]
            stack = stack[:l-2]
            value := v1 / v2
            stack = append(stack, value)
        } else {
            value, _ := strconv.Atoi(token)
            stack = append(stack, value)
        }
    }

    return stack[0]
}


