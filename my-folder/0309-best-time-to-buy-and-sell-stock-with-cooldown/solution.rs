impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut hold = -prices[0];
        let mut sold = 0;
        let mut rest = 0;

        for i in 0..prices.len() {
            let new_hold = hold.max(rest - prices[i]);
            let new_sold = hold + prices[i];
            let new_rest = rest.max(sold);

            hold = new_hold;
            sold = new_sold;
            rest = new_rest;
        }

        return rest.max(sold);
    }
}
