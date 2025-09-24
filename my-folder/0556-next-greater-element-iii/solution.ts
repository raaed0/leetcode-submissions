function nextGreaterElement(n: number): number {
    let m = [...n.toString()].map(d => parseInt(d))
    let len = m.length

    let i = len - 2
    while (i >= 0 && m[i] >= m[i+1])
        i--
    
    if (i < 0)
        return -1
    
        


    let j = len-1 
    while (m[j] <= m[i])
        j--
    

    [m[i], m[j]] = [m[j], m[i]]

    let l = i+1, r = len-1
    while (l<r) {
        [m[l], m[r]] = [m[r], m[l]]
        l++
        r--
    }

    let result = parseInt(m.join(""));
    return result > (2**31 - 1) ? -1 : result;
};
