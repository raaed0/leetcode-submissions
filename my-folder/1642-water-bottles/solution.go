func numWaterBottles(numBottles int, numExchange int) int {
    res := numBottles
    emptyBottles := numBottles
    
    for emptyBottles >= numExchange {
        tmp := emptyBottles / numExchange
        emptyBottles = emptyBottles - (tmp * numExchange) + tmp
        res += tmp
    }

    return res
}
