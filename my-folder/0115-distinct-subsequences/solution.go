func ok(i int, j int, s *string, t *string, dp *[][]int) int {
	n := len(*s)
	m := len(*t)

	if j == m {
		return 1
	}
	if i == n {
		return 0
	}

	if (*dp)[i][j] != -1 {
		return (*dp)[i][j]
	}

	total := 0
	if (*s)[i] == (*t)[j] {
		total = ok(i+1, j+1, s, t, dp) + ok(i+1, j, s, t, dp)
	} else {
		total = ok(i+1, j, s, t, dp)
	}

	(*dp)[i][j] = total
	return total
}

func numDistinct(s string, t string) int {
    var dp = make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(t)+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
	}

	return ok(0, 0, &s, &t, &dp)
}
