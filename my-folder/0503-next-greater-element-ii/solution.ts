class Stacks {
    private stack: number[]
    constructor() {
        this.stack = []
    }
    push(data: number) {
        this.stack.push(data)
    }
    pop(): number {
        if (this.isEmpty())
            throw new Error("Index out of bounds")
        return this.stack.pop()!
    }
    peak(): number {
        return this.stack[this.stack.length-1]
    }
    length(): number {
        return this.stack.length
    }
    isEmpty(): boolean {
        return this.length() == 0
    }
}

function nextGreaterElements(nums: number[]): number[] {
    let stack = new Stacks()
    let n = nums.length
    let result = new Array(n).fill(-1)

    for (let i = 0; i < 2 * n; i++) {
        let idx = i % n
        while (!stack.isEmpty() && nums[stack.peak()] < nums[idx]) {
            result[stack.pop()!] = nums[idx]
        }
        if (i < n) {
            stack.push(idx)
        } 

    }

    return result
};
