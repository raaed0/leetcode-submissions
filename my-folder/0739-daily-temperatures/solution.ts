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

function dailyTemperatures(temperatures: number[]): number[] {
    let stack = new Stacks()
    const n = temperatures.length
    let result: number[] = new Array(n).fill(0)

    for (let i = 0; i < temperatures.length; i++) {
        while (!stack.isEmpty() && temperatures[i] > temperatures[stack.peak()]) {
            const j = stack.pop()!
            result[j] = i - j
        }
        stack.push(i)
    }
    return result
};
