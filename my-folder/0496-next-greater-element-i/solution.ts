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

function nextGreaterElement(nums1: number[], nums2: number[]): number[] {
    let result: number[] = []
    let stack = new Stacks()
    let mapping = new Map<number, number>()

    for (let i = 0; i < nums2.length; i++) {
        if (stack.isEmpty()) {
            stack.push(nums2[i])
        } else if (stack.peak() >= nums2[i]) {
            stack.push(nums2[i])
        }
        // Recursion
        greaterNum(stack, nums2[i], mapping)
    }
    return nums1.map((v, i) => {
        return mapping.get(v) || -1
    })
}

function greaterNum(stack: Stacks, v: number, map: Map<number, number>) {
    if (stack.peak() < v) {
        let popped = stack.pop()
        map.set(popped, v)
        greaterNum(stack, v, map)
    }
    stack.push(v)
}

