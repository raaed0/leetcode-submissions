class Listnode<T> {
    val: T
    next: Listnode<T> | null
    constructor(val: T, next: Listnode<T> | null = null) {
        this.val = val
        this.next = next
    }
}

class SinglyLinkedList<T> {
    private head: Listnode<T> | null
    private tail: Listnode<T> | null
    private size: number
    constructor() {
        this.head = null
        this.tail = null
        this.size = 0
    }

    append(data: T): void {
        let newNode = new Listnode<T>(data)
        if (this.isEmpty()) {
            this.head = newNode
            this.tail = newNode
        } else {
            this.tail!.next = newNode
        }
        this.tail = newNode
        this.size++
    }

    removeTail(): T {
        if (this.isEmpty())
            throw new Error("Index out of bounds")

        let removedTail = this.tail
        if (this.head == this.tail) {
            this.head = null
            this.tail = null
            this.size--
            return removedTail!.val
        }

        let curr = this.head
        while (curr!.next != this.tail) {
            curr = curr!.next!
        }

        this.tail = curr
        this.size--
        return removedTail!.val
    }

    getTail(): T {
        if (!this.head)
            throw new Error("Index out of bounds")

        return this.tail!.val
    }

    isEmpty(): boolean {
        return !this.head
    }
}

class Stacks<T> {
    stack: SinglyLinkedList<T>
    constructor() {
        this.stack = new SinglyLinkedList<T>()
    }
    push(data: T) {
        this.stack.append(data)
    }
    pop(): T {
        return this.stack.removeTail()
    }
    peak(): T {
        return this.stack.getTail()
    }
    isEmpty(): boolean {
        return this.stack.isEmpty()
    }
}

function isValid(s: string): boolean {
    let stack = new Stacks<string>()
    let pairs: Record<string, string> = {
        "(": ")",
        "{": "}",
        "[": "]",
    }

    for (let i = 0; i < s.length; i++) {
        let char = s[i]
        if (Object.keys(pairs).includes(char)) {
            stack.push(char)
        } else if (Object.values(pairs).includes(char)) {
            if (stack.isEmpty())
                return false
            let popped = stack.pop()
            if (pairs[popped] != char){
                return false
            } else if (pairs[popped] == char) {
                continue
            }
        }
    }

    return stack.isEmpty()
};
