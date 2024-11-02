struct NumArray {
    nums: Vec<i32>,
    bit: Vec<i32>,
    n: usize,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NumArray {

    fn new(nums: Vec<i32>) -> Self {
        let mut num_array = NumArray {
            nums: nums.clone(),
            bit: vec![0; nums.len()+1],
            n: nums.len(),
        };

        for (i, &num) in nums.iter().enumerate() {
            num_array.update_bit(i + 1, num);
        };

        num_array
    }

    fn update_bit(&mut self, index: usize, delta: i32) {
        let mut idx = index;
        while idx <= self.n {
            self.bit[idx] += delta;
            idx += low_bit(idx);
        }
    }
    
    fn update(&mut self, index: i32, val: i32) {
        let delta = val - self.nums[index as usize];
        let mut idx = (index+1) as usize;

        self.nums[index as usize] += delta;
        while idx <= self.n {
            self.bit[idx] += delta;
            idx += low_bit(idx);
        }
    }
    
    fn sum_range(&self, left: i32, right: i32) -> i32 {
        self.prefix_sum(right+1) - self.prefix_sum(left)
    }

    fn prefix_sum(&self, index: i32) -> i32 {
        let mut sum = 0;
        let mut idx = index as usize;

        while idx > 0 {
            sum += self.bit[idx];
            idx -= low_bit(idx);
        }

        sum
    }
}

fn low_bit(i: usize) -> usize {
    i & (!i+1)
}

/**
 * Your NumArray object will be instantiated and called as such:
 * let obj = NumArray::new(nums);
 * obj.update(index, val);
 * let ret_2: i32 = obj.sum_range(left, right);
 */
