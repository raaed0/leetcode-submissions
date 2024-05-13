use std::collections::HashMap;
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut possible_pair_hash: HashMap<i32, i32> = HashMap::new();
        let mut sum_pair: Vec<i32> = Vec::new();

        for i in 0..nums.len() {
            let tmp = target - nums[i];

            if !possible_pair_hash.contains_key(&tmp) {
                possible_pair_hash.insert(nums[i], i as i32);
            } else {
                sum_pair.push(*possible_pair_hash.get(&tmp).unwrap());
                sum_pair.push(i as i32);
                return sum_pair;
            }
        }

        return vec![];
    }
}
