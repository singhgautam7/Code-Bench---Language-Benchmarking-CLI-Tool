use std::io::{self, Read};

fn quicksort(arr: &mut [i32]) {
    if arr.len() <= 1 { return; }
    let pivot = arr[arr.len() / 2];
    let mut left = 0;
    let mut right = arr.len() - 1;
    let mut i = 0;
    while i <= right {
        if arr[i] < pivot {
            arr.swap(i, left);
            left += 1;
            i += 1;
        } else if arr[i] > pivot {
            arr.swap(i, right);
            right -= 1;
        } else {
            i += 1;
        }
    }
    quicksort(&mut arr[0..left]);
    quicksort(&mut arr[right + 1..]);
}

fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).unwrap();
    let mut nums: Vec<i32> = input
        .split_whitespace()
        .filter_map(|s| s.parse().ok())
        .collect();
    if !nums.is_empty() {
        quicksort(&mut nums);
    }
}
