func restoreIpAddresses(s string) []string {
    res := []string{}
    if len(s) > 12 {
        return res
    }

    var dfs func(i, dots int, currIP string)
    dfs = func(i, dots int, currIP string) {
        if dots == 4 && i == len(s) {
            res = append(res, currIP[:len(currIP)-1])
            return
        }
        if dots > 4 {
            return
        }

        for j := i; j < min(i+3, len(s)); j++ {
            if i != j && s[i] == '0' {
                continue
            }
            part := s[i : j+1]
            num, _ := strconv.Atoi(part)
            if num < 256 {
                dfs(j+1, dots+1, currIP+part+".")
            }
        }
    }

    dfs(0, 0, "")
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
