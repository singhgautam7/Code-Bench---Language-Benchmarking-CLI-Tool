use std::io::{self, Read};

/// Perform binary search on a sorted slice.
/// Returns the index of target, or -1 if not found.
fn binary_search(arr: &[i64], target: i64) -> i64 {
    let mut low: i64 = 0;
    let mut high: i64 = arr.len() as i64 - 1;

    while low <= high {
        let mid = low + (high - low) / 2;
        if arr[mid as usize] == target {
            return mid;
        } else if arr[mid as usize] < target {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }

    -1
}

fn main() {
    let mut input = String::new();
    io::stdin().read_to_string(&mut input).expect("Failed to read stdin");

    let tokens: Vec<i64> = input
        .split_whitespace()
        .map(|s| s.parse::<i64>().expect("Failed to parse integer"))
        .collect();

    if tokens.len() < 2 {
        eprintln!("Error: expected array and target on stdin");
        std::process::exit(1);
    }

    let target = tokens[tokens.len() - 1];
    let arr = &tokens[..tokens.len() - 1];

    let result = binary_search(arr, target);
    println!("{}", result);
}
