func maxProfit(prices []int) int {
    profit := 0
    bought := prices[0]

    for _, v := range prices {
        bought = min(bought, v)
        profit = max(profit, v - bought)
    }

    return profit
}
