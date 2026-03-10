impl Solution {
    pub fn max_profit(k: i32, prices: Vec<i32>) -> i32 {
        let k = k as usize;
        let mut buy: Vec<i32> = vec![i32::MIN; k+1];
        let mut sell: Vec<i32> = vec![0; k+1];

        for price in prices {
            for i in 1..=k {
                buy[i] = buy[i].max(sell[i-1] - price);
                sell[i] = sell[i].max(buy[i] + price);
            }
        }
        return sell[k];        
    }
}
