function leastInterval(tasks: string[], n: number): number {
    let taskCount = new Array(26).fill(0)

    for (let task of tasks) {
        taskCount[task.charCodeAt(0) - 'A'.charCodeAt(0)]++ 
    }

    let maxHeap = new MaxPriorityQueue()
    for (let i = 0; i < 26; i++) {
        if (taskCount[i] > 0) {
            maxHeap.push(taskCount[i])
        }
    }

    let time = 0
    let q = new Queue()

    while (maxHeap.size() > 0 || q.size() > 0) {
        time++

        if (maxHeap.size() > 0) {
            let tmp: number = (maxHeap.pop() as number)-1
            if (tmp !== 0) {
                q.push([tmp, time+n])
            }
        }

        if (q.size() > 0 && q.front()[1] === time) {
            maxHeap.push(q.pop()[0])
        }
    }

    return time
};
