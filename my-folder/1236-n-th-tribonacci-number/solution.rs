impl Solution {
    pub fn tribonacci(n: i32) -> i32 {
        let mut cache: Vec<i32> = vec![0, 1, 1];
        let n = n as usize;
        if n < 3 {
            return cache[n];
        }

        for i in 3..=n {
            cache[i % 3] = cache[0] + cache[1] + cache[2];
        }

        return cache[n % 3];
    }
}
