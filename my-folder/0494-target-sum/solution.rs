use std::collections::HashMap;

impl Solution {
    pub fn find_target_sum_ways(nums: Vec<i32>, target: i32) -> i32 {
        let mut dp: HashMap<i32, i32> = HashMap::new();
        dp.insert(0, 1);

        for &num in &nums {
            let mut next_dp: HashMap<i32, i32> = HashMap::new();

            for (&total, &count) in &dp {
                *next_dp.entry(total + num).or_insert(0) += count;
                *next_dp.entry(total - num).or_insert(0) += count;
            }
            dp = next_dp;
        }

        *dp.get(&target).unwrap_or(&0)
    }
}
