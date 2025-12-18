func calPoints(operations []string) int {
    stack := make([]int, 0, len(operations))
    sum := 0

    for _, op:= range operations {
        switch op {
            case "+": 
                a := stack[len(stack) - 1]
                b := stack[len(stack) - 2]
                v := a + b
                stack = append(stack, v)
                sum += v
            case "D":
                v := 2 * stack[len(stack) - 1]
                stack = append(stack, v)
                sum += v
            case "C":
                v := stack[len(stack) - 1]
                stack = stack[:len(stack)-1]
                sum -= v
            default:
                v, _ := strconv.Atoi(op)
                stack = append(stack, v)
                sum += v
        }
    }

    return sum
}
