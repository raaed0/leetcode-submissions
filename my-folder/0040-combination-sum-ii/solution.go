func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    var res [][]int
    backtrack([]int{}, target, 0, candidates, &res)
    return res
}

func backtrack(path []int, remain int, start int, candidates []int, res *[][]int) {
    if remain == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *res = append(*res, tmp)
        return
    }

    if remain < 0 {
        return
    }

    for i := start; i<len(candidates); i++ {
        if i > start && candidates[i] == candidates[i-1] {
            continue
        }

        if candidates[i] > remain {
            break
        }

        backtrack(append(path, candidates[i]), remain-candidates[i], i+1, candidates, res)
    }
}
