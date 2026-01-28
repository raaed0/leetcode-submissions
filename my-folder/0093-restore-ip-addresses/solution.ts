function restoreIpAddresses(s: string): string[] {
    let res: string[] = []

    if (s.length > 12) {
        return res
    }

    let backtracking = function(i: number, dots: number, currIP: string) {
        if (dots === 4 && i === s.length) {
            res.push(currIP.slice(0, currIP.length-1))
            return
        }

        if (dots > 4)
            return

        for (let j = i; j < Math.min(i+3, s.length); j++) {
            if (i !== j && s[i] === "0") {
                continue
            }
            let num = parseInt(s.slice(i, j+1))
            if (num < 256) {
                backtracking(j+1, dots+1, currIP+num+".")
            }
        }
    }

    backtracking(0, 0, "")
    return res
};
