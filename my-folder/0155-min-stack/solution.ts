class MinStack {
    private stack: number[]
    private minValue: number[] = []

    constructor() {
        this.stack = []
    }

    push(val: number): void {
        this.stack.push(val)

        if (this.minValue.length == 0 || val <= this.minValue[this.minValue.length - 1])
            this.minValue.push(val)
    }

    pop(): void {
        let popped = this.stack.pop()

        if (popped == this.minValue[this.minValue.length -1 ])
            this.minValue.pop()
    }

    top(): number {
        return this.stack[this.stack.length-1]
    }

    getMin(): number {
        return this.minValue[this.minValue.length - 1]
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = new MinStack()
 * obj.push(val)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */
