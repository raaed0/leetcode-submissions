/**
 Do not return anything, modify s in-place instead.
 */
function reverseString(s: string[]): void {
    let head: number = 0
    let tail: number = s.length-1

    while (head < tail) {
        let tmp = s[head]
        s[head] = s[tail]
        s[tail] = tmp

        head++
        tail--
    }
};
