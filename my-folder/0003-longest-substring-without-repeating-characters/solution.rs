use std::collections::HashMap;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
    let mut char_index_map: HashMap<char, usize> = HashMap::new();
    let mut longest = 0;
    let mut start = 0;

    for (i, char) in s.chars().enumerate() {
        if let Some(prev) = char_index_map.insert(char, i) {
            start = start.max(prev + 1);
        }
        longest = longest.max(i - start + 1)
    }
    return longest as i32; 
    }
}
