function combinationSum2(candidates: number[], target: number): number[][] {
    candidates.sort()
    let result = []
    backtrack(0, [], candidates, 0, target, result)
    return result
};

function backtrack(i: number, path: number[], candidates: number[], total: number, target: number, result: number[][]) {
    if (total === target) {
        result.push([...path])
        return
    }

    if (i >= candidates.length || total > target)
        return

    path.push(candidates[i])
    backtrack(i+1, path, candidates, total+candidates[i], target, result)

    while (i+1 < candidates.length && candidates[i]==candidates[i+1])
        i++

    path.pop()
    backtrack(i+1, path, candidates, total, target, result)
}
